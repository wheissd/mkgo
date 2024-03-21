package main

import (
	"errors"
	"fmt"
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

type runInfo struct {
	EntHash string `json:"ent_hash"`
}

const runInfoFile = "mkgo/run_info.json"

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

	runInfoBytes, err := os.ReadFile(runInfoFile)
	if err != nil {
		return err
	}
	var ri runInfo
	if err = json.Unmarshal(runInfoBytes, &ri); err != nil {
		return err
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

	if ri.EntHash != hash {
		cmd.logger.Debug("generate ent")
		// gen ent
		entPath := "./ent/cmd"
		if !internalMode {
			entPath = "./internal/ent/cmd"
		}
		err = cmd.runCmd("go run " + entPath)
		if err != nil {
			return err
		}

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
			err = cmd.pre(preOutput, schemaPath, genPath)
			if err != nil {
				return err
			}
		}

		err = cmd.runCmd("atlas migrate hash --dir file://internal/ent/migrate/migrations")
		if err != nil {
			return err
		}

		err = cmd.runCmd("atlas migrate diff " +
			"--dir file://internal/ent/migrate/migrations " +
			"--to ent://internal/ent/schema " +
			"--schema dev " +
			"--dev-url docker://postgres/15/test?search_path=public")
		if err != nil {
			return err
		}

		ri.EntHash = hash
	}

	cfgDir := "internal/config"
	if internalMode {
		cfgDir = "config"
	}

	cmd.logGreen("genConfig")
	cmd.logger.Debug("genConfig", zap.String("target", cfgDir))
	cmd.genConfig(cfgDir)

	var apiCfg config.GenConfigItem
	for i := range cfg.APIs {
		apiCfg = cfg.APIs[i]
		break
	}

	apigenPath := fmt.Sprintf("%s/%s%s%s/cmd/apigen", pkgInfo.Pkg, rootDir, apigenDelim, apiCfg.OutputPath)
	err = cmd.runCmd("go run " + apigenPath + " -ent_path=" + pkgInfo.RootDir + "/ent/gen")
	if err != nil {
		return err
	}

	for _, apiCfg := range cfg.APIs {
		if apiCfg.Transport == gen.TransportHTTP || apiCfg.Transport == gen.TransportDefault {
			ogenTarget := rootDir + apigenDelim + apiCfg.OutputPath + "/ogen"
			//if !internalMode {
			//	ogenTarget = "internal/" + ogenTarget
			//}
			err = cmd.runCmd("go run github.com/ogen-go/ogen/cmd/ogen@latest --package " +
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
			err = cmd.runCmd("protoc --proto_path proto --go_out=. --go-grpc_out=. --go-grpc_opt=paths=import " + strings.Join(protoFiles, " "))
			if err != nil {
				return err
			}
		}
	}

	if runInfoBytes, err = json.Marshal(ri); err != nil {
		return err
	}
	cmd.logger.Debug("runInfo write",
		zap.String("path", runInfoFile),
		zap.ByteString("runInfo", runInfoBytes),
	)
	if err = os.WriteFile(runInfoFile, runInfoBytes, 0744); err != nil {
		return err
	}
	cmd.logger.Debug("runInfo write success", zap.String("path", runInfoFile))

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

func (cmd *cmd) logGreen(s string) {
	cmd.logger.Info("\u001b[32m" + s + "\u001b[0m")
}

func (cmd *cmd) runCmd(cmdStr string) error {
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
			return e
		}
		cmd.logger.Error("runCmd err", zap.Error(err), zap.ByteString("output", output))
		return err
	}
	fmt.Print(string(output))
	return nil
}
