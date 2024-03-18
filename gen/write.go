package gen

import (
	"bytes"
	"embed"
	"fmt"
	"go/format"
	"os"
	"os/exec"
	"strings"
	"text/template"

	"github.com/wheissd/mkgo/internal/entity"
	"go.uber.org/zap"
)

type templateConfig struct {
	file       string
	withFiles  []string
	template   *template.Template
	outputFile string
	transport  string
}

var (
	//go:embed templates/service/*.gotmpl
	//go:embed templates/ogen/*.gotmpl
	//go:embed templates/grpc/*.gotmpl
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
