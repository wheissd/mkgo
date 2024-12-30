package main

import (
	"bytes"
	"embed"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path"
	"sort"
	"strings"
	"text/template"

	"github.com/wheissd/mkgo/internal/config"
	"github.com/wheissd/mkgo/internal/parse"
	"go.uber.org/zap"
)

var (
	//go:embed templates/*.gotmpl
	//go:embed templates/project/internal/api/http/spec/*.gotmpl
	//go:embed templates/project/internal/api/http/handler/*.gotmpl
	preTmpl embed.FS
)

type Entity struct {
	Schema     string
	Model      string
	FieldOrder map[string]int
}

type PreSchema struct {
	SchemaPath string
	GenPath    string
	Entities   []Entity
	Path       string
}

type PreMainSchema struct {
	Module    string
	Cfg       config.GenConfigItem
	IsGRPC    bool
	IsOpenAPI bool
	RootDir   string
}

func (cmd *cmd) pre(output, schemaPath, genPath, rootDir string, cfg config.GenConfigItem) error {

	wd, _ := os.Getwd()
	cmd.logger.Debug("working dir: " + wd)

	module, level := parse.GetModLevel(wd, 0)
	wdLevel := strings.Repeat("../", level)

	schemaImport := module + "/" + schemaPath
	genImport := module + "/" + genPath

	ts := []struct {
		template string
		path     string
		name     string
	}{
		{
			template: "templates/pre_main.gotmpl",
			path:     output + "/main.go",
			name:     "apigen main",
		},
		{
			template: "templates/project/internal/api/http/spec/spec.gotmpl",
			path:     fmt.Sprintf("%s/%s/http/spec/%s", rootDir, cfg.OutputPath, "/spec.go"),
			name:     "http spec",
		},
		{
			template: "templates/project/internal/api/http/handler/handler.gotmpl",
			path:     fmt.Sprintf("%s/%s/http/handler/%s", rootDir, cfg.OutputPath, "/handler.go"),
			name:     "http handler",
		},
	}

	schema := PreMainSchema{
		Module:    module,
		Cfg:       cfg,
		IsGRPC:    cfg.ProtoPath != "",
		IsOpenAPI: cfg.OpenApiPath != "",
		RootDir:   rootDir,
	}

	for _, t := range ts {
		tmpl, err := template.ParseFS(preTmpl, t.template)
		if err != nil {
			panic(err)
		}

		if _, err = os.Stat(t.path); err != nil {
			cmd.logger.Debug(fmt.Sprintf("empty %s, add %s", t.name, t.path))
			buf := bytes.NewBuffer(nil)
			if err = tmpl.Execute(buf, schema); err != nil {
				panic(err)
			}

			b, err := format.Source(buf.Bytes())
			if err != nil {
				cmd.logger.Error("format.Source(buf.Bytes())", zap.Error(err))

				return err
			}

			dirPath := path.Dir(t.path)

			// Create the directories if they don't exist
			err = os.MkdirAll(dirPath, os.ModePerm)
			if err != nil {
				panic(err)
			}

			err = os.WriteFile(t.path, b, 0744)
			if err != nil {
				cmd.logger.Error(fmt.Sprintf("os.WriteFile(\"%s\", b, 0744)", t.path), zap.Error(err))
				return err
			}
		}
	}

	preTemplate, err := template.ParseFS(preTmpl, "templates/pre_gen.gotmpl")
	if err != nil {
		panic(err)
	}

	schemaDir := wdLevel + schemaPath
	cmd.logger.Debug(
		"parser.ParseDir",
		zap.String("wd", wd),
		zap.String("schemaImport", schemaImport),
		zap.String("wdLevel", wdLevel),
		zap.String("schemaDir", schemaDir),
		zap.String("output", output),
		zap.String("schemaPath", schemaPath),
		zap.String("genPath", genPath),
	)

	// parse schema files
	schemaPkgs, err := parser.ParseDir(
		token.NewFileSet(),
		schemaDir,
		func(_ fs.FileInfo) bool { return true },
		0,
	)
	if err != nil {
		cmd.logger.Error("parser.ParseDir Schema", zap.Error(err))
		return err
	}

	cmd.logger.Debug("output Path: " + output)

	sch := PreSchema{
		SchemaPath: schemaImport,
		GenPath:    genImport,
		Entities:   make([]Entity, 0),
		Path:       output,
	}
	for pkgIndex := range schemaPkgs {
		for _, file := range schemaPkgs[pkgIndex].Files {
			for _, decl := range file.Decls {
				if typeDecl, ok := decl.(*ast.GenDecl); ok {
					if typeDecl.Tok == token.TYPE {
						cmd.logger.Debug("entity: " + typeDecl.Specs[0].(*ast.TypeSpec).Name.Name)
						sch.Entities = append(sch.Entities, Entity{
							Schema: typeDecl.Specs[0].(*ast.TypeSpec).Name.Name,
							Model:  "gen." + typeDecl.Specs[0].(*ast.TypeSpec).Name.Name,
						})
					}
				}
			}
		}
	}

	sort.Slice(sch.Entities, func(i, j int) bool {
		return sch.Entities[i].Schema < sch.Entities[j].Schema
	})

	buf := bytes.NewBuffer(nil)

	err = preTemplate.Execute(buf, sch)
	if err != nil {
		cmd.logger.Error("preTemplate.Execute(buf, sch)", zap.Error(err))
		return err
	}

	b, err := format.Source(buf.Bytes())
	if err != nil {
		cmd.logger.Error("format.Source(buf.Bytes())", zap.Error(err))

		return err
	}

	cmd.logger.Debug("output file: " + sch.Path + "/pre_gen.go")
	_ = os.Remove(sch.Path + "/pre_gen.go")

	err = os.WriteFile(sch.Path+"/pre_gen.go", b, 0744)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
