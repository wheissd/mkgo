package main

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"strings"

	"github.com/ogen-go/ogen/json"
	"github.com/rogpeppe/go-internal/dirhash"
	"github.com/urfave/cli/v2"
	"github.com/wheissd/mkgo/gen"
	config2 "github.com/wheissd/mkgo/gen/config"
	"github.com/wheissd/mkgo/internal/config"
	"github.com/wheissd/mkgo/internal/parse"
	"go.uber.org/zap"
)

func (cmd *cmd) generate(ctx *cli.Context) error {
	var cfg config.GenConfig
	config2.Parse(&cfg, config2.OptionPath("mkgo_config.yaml"))
	pkgInfo := parse.GetPkgInfo()

	_, err := os.ReadDir("internal")
	var (
		internalMode = false
	)
	if err != nil {
		internalMode = true
	}

	cmd.logger.Debug("start reading runInfoFile")
	runInfoBytes, err := os.ReadFile(runInfoFile)
	unmarshalRunInfo := true
	if err != nil {
		if err != os.ErrNotExist && !errors.Is(err, fs.ErrNotExist) {
			cmd.logger.Error("start reading runInfoFile err", zap.Error(err))
			return err
		}
		unmarshalRunInfo = false
	}
	var ri runInfo
	if unmarshalRunInfo {
		if err = json.Unmarshal(runInfoBytes, &ri); err != nil {
			return err
		}
	}

	if ctx.Value("skipDepCheck") == nil || ctx.Value("skipDepCheck") == false {
		err := cmd.checkDependencies(&ri)
		if err != nil {
			cmd.logger.Error("cmd.checkDependencies()", zap.Error(err))
		}
	}

	entPath := "./internal/ent"
	if internalMode {
		entPath = "./ent"
	}
	hash, err := dirhash.HashDir(entPath, "mkgo_hash", dirhash.Hash1)
	if err != nil {
		return err
	}
	cmd.logger.Debug("mkgo_hash", zap.String("prev", ri.EntHash), zap.String("current", hash))

	rootDir := pkgInfo.RootDir
	apigenDelim := ""
	if rootDir == "" {
		apigenDelim = "/"
		rootDir = "internal"
	}
	if internalMode {
		apigenDelim = "/"
	}

	cfgHash, err := dirhash.Hash1([]string{"mkgo_config.yaml"}, func(name string) (io.ReadCloser, error) {
		return os.Open("mkgo_config.yaml")
	})
	if err != nil {
		return err
	}

	if !ri.entHashCheck(hash) {
		cmd.logger.Debug("generate ent")
		// gen ent
		entPath := "./ent/cmd"
		if !internalMode {
			entPath = "./internal/ent/cmd"
		}
		err = cmd.runVCmd("go run " + entPath)
		if err != nil {
			return err
		}

		err = cmd.runVCmd("atlas migrate hash --dir file://internal/ent/migrate/migrations")
		if err != nil {
			return err
		}

		err = cmd.runVCmd("atlas migrate diff " +
			"--dir file://internal/ent/migrate/migrations " +
			"--to ent://internal/ent/schema " +
			"--schema dev " +
			"--dev-url docker://postgres/15/test?search_path=public")
		if err != nil {
			return err
		}
	}
	if !ri.hashCheck(hash, cfgHash) {
		// gen pre
		for _, cfgItem := range cfg.APIs {
			preOutput := "./" + rootDir + apigenDelim + cfgItem.OutputPath + "/cmd/apigen"
			schemaPath := rootDir + "/ent/schema"
			genPath := rootDir + "/ent/gen"
			cmd.logGreen("run cmd.pre")
			cmd.logger.Debug(
				"run cmd.pre",
				zap.String("preOutput", preOutput),
				zap.String("rootDir", rootDir),
				zap.String("schemaPath", schemaPath),
				zap.String("genPath", genPath),
			)
			err = cmd.pre(preOutput, schemaPath, genPath, rootDir, cfgItem)
			if err != nil {
				return err
			}
		}
	}
	ri.EntHash = hash
	ri.CfgHash = cfgHash

	cfgDir := "internal/config"
	if internalMode {
		cfgDir = "config"
	}

	cmd.logGreen("genConfig")
	cmd.logger.Debug("genConfig", zap.String("target", cfgDir))
	cmd.genConfig(cfgDir)

	for _, apiCfg := range cfg.APIs {
		apigenPath := fmt.Sprintf("%s/%s%s%s/cmd/apigen", pkgInfo.Pkg, rootDir, apigenDelim, apiCfg.OutputPath)
		err = cmd.runVCmd("go run " + apigenPath + " -ent_path=" + pkgInfo.RootDir + "/ent" + " -mode=" + apiCfg.Mode)
		if err != nil {
			return err
		}

		if apiCfg.Transport == gen.TransportHTTP || apiCfg.Transport == gen.TransportDefault {
			ogenTarget := rootDir + apigenDelim + apiCfg.OutputPath + "/ogen"
			//if !internalMode {
			//	ogenTarget = "internal/" + ogenTarget
			//}
			err = cmd.runVCmd("ogen --package " +
				"ogen --target " + ogenTarget + " --clean openapi/" + apiCfg.Mode + ".json --convenient-errors on")
			if err != nil {
				return err
			}
		} else {
			err = cleanGrpcDir(fmt.Sprintf("%s%s", rootDir, apigenDelim), apiCfg)
			if err != nil {
				return err
			}
			files, err := os.ReadDir("proto")
			if err != nil {
				return err
			}
			protoFiles := []string{}
			for _, file := range files {
				if strings.HasSuffix(file.Name(), ".proto") {
					protoFiles = append(protoFiles, file.Name())
				}
			}
			err = cmd.runVCmd("protoc --proto_path proto --go_out=. --go-grpc_out=. --go-grpc_opt=paths=import " + strings.Join(protoFiles, " "))
			if err != nil {
				return err
			}
		}
	}

	if runInfoBytes, err = json.Marshal(ri); err != nil {
		return err
	}
	cmd.logger.Debug("runInfo write",
		zap.String("Path", runInfoFile),
		zap.ByteString("runInfo", runInfoBytes),
	)
	if err = os.WriteFile(runInfoFile, runInfoBytes, 0744); err != nil {
		return err
	}
	cmd.logger.Debug("runInfo write success", zap.String("Path", runInfoFile))

	return nil
}

func cleanGrpcDir(rootDir string, cfg config.GenConfigItem) error {
	path := rootDir + cfg.OutputPath + "/grpc"
	dir, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for _, f := range dir {
		if strings.HasSuffix(f.Name(), "pb.go") {
			err = os.Remove(path + "/" + f.Name())
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (cmd *cmd) logRed(s string) {
	cmd.logger.Info("\u001b[32m" + s + "\u001b[0m")
}

func (cmd *cmd) logGreen(s string) {
	cmd.logger.Info("\u001b[32m" + s + "\u001b[0m")
}

// runVCmd v As void
func (cmd *cmd) runVCmd(cmdStr string) error {
	output, err := cmd.runCmd(cmdStr)
	if err != nil {
		return err
	}
	fmt.Print(string(output))
	return nil
}

func (cmd *cmd) runCmd(cmdStr string) ([]byte, error) {
	cmd.logGreen(cmdStr)
	cmdSplitted := strings.Split(cmdStr, " ")
	c := exec.Command(cmdSplitted[0], cmdSplitted[1:]...)
	output, err := c.Output()
	if err != nil {
		var e *exec.ExitError
		if errors.As(err, &e) {
			//cmd.logger.Debug("run cmd err")
			if strings.Contains(cmdStr, "apigen") {
				cmd.logger.Error("run cmd err", zap.Error(err))
			}
			fmt.Println(string(e.Stderr))
			return nil, e
		}
		cmd.logger.Error("runVCmd err", zap.Error(err), zap.ByteString("output", output))
		return nil, err
	}
	return output, nil
}
