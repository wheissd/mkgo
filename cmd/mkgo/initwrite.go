package main

import (
	"bytes"
	"embed"
	"go/format"
	"os"
	"strings"
	"text/template"
)

type templateConfig struct {
	file       string
	template   *template.Template
	outputFile string
}

var (
	//go:embed templates/project/internal/app/*.gotmpl
	//go:embed templates/project/internal/config/*tmpl
	//go:embed templates/project/internal/db/*tmpl
	//go:embed templates/project/internal/ent/*.gotmpl
	//go:embed templates/project/internal/ent/cmd/*.gotmpl
	//go:embed templates/project/internal/ent/schema/*.gotmpl
	//go:embed templates/project/*tmpl
	//go:embed templates/project/mkgo/*tmpl
	tmpl embed.FS
)

type Tool struct {
	// full tool module path
	Module string
}

type Schema struct {
	// module path
	Path string
	// project path
	ProjectPath string
	// module from root go.mod
	Module string
	// target module name
	TargetName   string
	Tool         Tool
	InternalMode bool
}

func writeGoFiles(sch Schema) error {
	cfg := []*templateConfig{
		{
			file: "generate.gotmpl",
		},
		{
			file: "config.yamltmpl",
		},
		{
			file: ".gitignoretmpl",
		},
		//{
		//	file: "cmd/restclient/main.gotmpl",
		//},
		//{
		//	file: "cmd/grpcadmin/main.gotmpl",
		//},
		{
			file: "internal/app/di.gotmpl",
		},
		{
			file: "internal/db/db.gotmpl",
		},
		//{
		//	file: "internal/grpc_admin/cmd/apigen/main.gotmpl",
		//},
		//{
		//	file: "internal/grpc_admin/cmd/apigen/pre_gen.gotmpl",
		//},
		//{
		//	file: "internal/grpc_admin/grpc/module.gotmpl",
		//},
		//{
		//	file: "internal/rest_client/cmd/apigen/main.gotmpl",
		//},
		//{
		//	file: "internal/rest_client/cmd/apigen/pre_gen.gotmpl",
		//},
		//{
		//	file: "internal/rest_client/http/server.gotmpl",
		//},
		//{
		//	file: "internal/rest_client/grpc/module.gotmpl",
		//},
		//{
		//	file: "internal/rest_client/http/handler/handler.gotmpl",
		//},
		{
			file: "internal/ent/cmd/main.gotmpl",
		},
		{
			file: "internal/ent/withtx.gotmpl",
		},
		{
			file: "internal/ent/schema/default_model.gotmpl",
		},
		{
			file: "mkgo_config.yamltmpl",
		},
		{
			file: "mkgo/fields_order_gen.jsontmpl",
		},
		{
			file: "mkgo/run_info.jsontmpl",
		},
		{
			file: "mkgo/fields_order_gen.jsontmpl",
		},
		{
			file: "internal/config/config.gotmpl",
		},
		{
			file: "internal/config/gen.gotmpl",
		},
	}

	// read templates
	for _, cfgItem := range cfg {
		dirSplitted := strings.Split(cfgItem.file, "/")
		var fName, fPath string
		if len(dirSplitted) > 1 {
			fName = dirSplitted[len(dirSplitted)-1]
			fPath = "templates/project/" + strings.Join(dirSplitted[:len(dirSplitted)], "/")
		} else {
			fName = cfgItem.file
			fPath = "templates/project/" + fName
		}
		var err error
		if cfgItem.template, err = template.New(fName).
			ParseFS(tmpl, fPath); err != nil {
			return err
		}

		pathContainsInternal := strings.Contains(sch.Path, "internal")
		oPath := strings.ReplaceAll(
			strings.ReplaceAll(fPath, "tmpl", ""),
			"templates/project/",
			"")
		if pathContainsInternal {
			oPath = strings.ReplaceAll(oPath, "internal/", "")
		}
		cfgItem.outputFile = sch.ProjectPath + oPath
	}
	for _, t := range cfg {
		buf := bytes.NewBuffer(nil)

		err := t.template.Execute(buf, sch)
		if err != nil {
			panic(err)
		}
		var b = buf.Bytes()
		if strings.HasSuffix(t.outputFile, "go") {
			b, err = format.Source(buf.Bytes())
			if err != nil {
				panic(err)
			}
		}

		_ = os.Remove(t.outputFile)

		err = os.WriteFile(t.outputFile, b, 0744)
		if err != nil {
			panic(err)
		}
	}
	return nil
}
