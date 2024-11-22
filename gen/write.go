package gen

import (
	"bytes"
	"embed"
	"fmt"
	"go/format"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
	"text/template"

	"github.com/wheissd/mkgo/internal/entity"
	"go.uber.org/zap"
)

type writeContext struct {
	path string
}

type whenWrite func(ctx writeContext) bool

type templateConfig struct {
	file       string
	withFiles  []string
	template   *template.Template
	outputFile string
	transport  string
	when       whenWrite
}

func whenPathNotExist(ctx writeContext) bool {
	_, err := os.Stat(ctx.path)
	return err == nil
}

var (
	//go:embed templates/service/*.gotmpl
	//go:embed templates/ogen/*.gotmpl
	//go:embed templates/grpc/*.gotmpl
	//go:embed templates/http/*.gotmpl
	//go:embed templates/http/handler/*.gotmpl
	tmpl embed.FS
)

func writeGoFiles(logger *zap.Logger, sch *entity.Schema) {
	cfg := []*templateConfig{
		{
			file: "service/service.gotmpl",
			withFiles: []string{
				"service/edges.gotmpl",
			},
		},
		{
			file: "service/modifier.gotmpl",
		},
		{
			file: "service/entities.gotmpl",
		},
		{
			file: "service/query.gotmpl",
		},
		{
			file:      "ogen/handler.gotmpl",
			transport: TransportHTTP,
		},
		{
			file:      "ogen/responses.gotmpl",
			transport: TransportHTTP,
		},
		{
			file:      "ogen/edges.gotmpl",
			transport: TransportHTTP,
		},
		{
			file:      "grpc/service_impl_gen.gotmpl",
			transport: TransportGRPC,
		},
		{
			file:      "grpc/models_gen.gotmpl",
			transport: TransportGRPC,
		},
		{
			file:      "grpc/module.gotmpl",
			transport: TransportGRPC,
		},
		{
			file:      "http/server.gotmpl",
			when:      whenPathNotExist,
			transport: TransportHTTP,
		},
		{
			file:      "http/handler/handler.gotmpl",
			when:      whenPathNotExist,
			transport: TransportHTTP,
		},
	}
	_, err := os.ReadDir("internal")
	var (
		internalMode = false
	)
	if err != nil {
		internalMode = true
	}
	for _, cfgItem := range cfg {
		if cfgItem.transport != "" && cfgItem.transport != sch.Cfg.Transport {
			continue
		}
		dirSplitted := strings.Split(cfgItem.file, "/")
		var fName, fPath string
		if len(dirSplitted) > 1 {
			fName = dirSplitted[len(dirSplitted)-1]
			fPath = "templates/" + strings.Join(dirSplitted[:len(dirSplitted)], "/")
		} else {
			fName = cfgItem.file
			fPath = "templates/" + fName
		}

		patterns := []string{fPath}
		for _, p := range cfgItem.withFiles {
			patterns = append(patterns, "templates/"+p)
		}
		var err error
		if cfgItem.template, err = template.New(fName).Funcs(fns).
			ParseFS(tmpl, patterns...); err != nil {
			panic(err)
		}

		oPath := strings.ReplaceAll(
			strings.ReplaceAll(fPath, "tmpl", ""),
			"templates/",
			"",
		)

		outputFile := sch.RootPkg + sch.Cfg.OutputPath + "/" + oPath
		if internalMode {
			outputFile = sch.Cfg.OutputPath + "/" + oPath
		}

		if cfgItem.when != nil && cfgItem.when(writeContext{
			path: outputFile,
		}) {
			logger.Warn(fmt.Sprintf("skip writing file. reason: %s", reflect.TypeOf(cfgItem.outputFile).Name()))
			continue
		}

		buf := bytes.NewBuffer(nil)

		err = cfgItem.template.Execute(buf, sch)
		if err != nil {
			panic(err)
		}
		var b = buf.Bytes()
		if strings.HasSuffix(cfgItem.outputFile, "go") {
			b, err = format.Source(buf.Bytes())
			if err != nil {
				panic(err)
			}
		}

		//_ = os.Remove(cfgItem.outputFile)

		err = createMissingDirs(outputFile)
		if err != nil {
			panic(err)
		} else {
			logger.Debug(fmt.Sprintf("createMissingDirs: %s", outputFile))
		}

		logger.Debug("writing file: " + outputFile)
		err = os.WriteFile(outputFile, b, 0744)
		if err != nil {
			panic(err)
		}
	}

	importsPath := sch.RootPkg + sch.Cfg.OutputPath
	if internalMode {
		importsPath = sch.Cfg.OutputPath
	}

	runCmd("goimports -l -w " + importsPath)
}

// createMissingDirs creates only missing directories from the given file path.
func createMissingDirs(filePath string) error {
	// Get the directory part of the file path
	dir := filepath.Dir(filePath)

	// Create the directories with necessary permissions (0755 by default)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directories for path %s: %w", filePath, err)
	}

	return nil
}

func runCmd(cmdStr string) error {
	fmt.Println("\u001b[32m" + cmdStr + "\u001b[0m")
	cmdSplitted := strings.Split(cmdStr, " ")
	cmd := exec.Command(cmdSplitted[0], cmdSplitted[1:]...)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(output)
		return err
	}
	fmt.Print(string(output))
	return nil
}
