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

	// @TODO remove duplicate code
	mainTemplate, err := template.ParseFS(preTmpl, "templates/pre_main.gotmpl")
	if err != nil {
		panic(err)
	}

	mainSchema := PreMainSchema{
		Module:    module,
		Cfg:       cfg,
		IsGRPC:    cfg.ProtoPath != "",
		IsOpenAPI: cfg.OpenApiPath != "",
		RootDir:   rootDir,
	}
	if _, err = os.Stat(output + "/main.go"); err != nil {
		cmd.logger.Debug("empty apigen main, add " + output + "/main.go")
		buf := bytes.NewBuffer(nil)
		if err = mainTemplate.Execute(buf, mainSchema); err != nil {
			panic(err)
		}

		b, err := format.Source(buf.Bytes())
		if err != nil {
			cmd.logger.Error("format.Source(buf.Bytes())", zap.Error(err))

			return err
		}

		err = os.WriteFile(output+"/main.go", b, 0744)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	// @TODO remove duplicate code
	specTemplate, err := template.ParseFS(preTmpl, "templates/project/internal/api/http/spec/spec.gotmpl")
	if err != nil {
		panic(err)
	}

	specDirPath := fmt.Sprintf("%s/%s/http/spec", rootDir, cfg.OutputPath)
	specPath := fmt.Sprintf("%s/%s", specDirPath, "/spec.go")
	if _, err = os.Stat(specPath); cfg.OpenApiPath != "" && err != nil {
		cmd.logger.Debug("empty spec, add " + specPath)
		buf := bytes.NewBuffer(nil)
		if err = specTemplate.Execute(buf, nil); err != nil {
			panic(err)
		}

		b, err := format.Source(buf.Bytes())
		if err != nil {
			cmd.logger.Error("format.Source(buf.Bytes())", zap.Error(err))

			return err
		}

		// Create the directories if they don't exist
		err = os.MkdirAll(specDirPath, os.ModePerm)
		if err != nil {
			panic(err)
		}

		err = os.WriteFile(specPath, b, 0744)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	// @TODO remove duplicate code
	handlerTemplate, err := template.ParseFS(preTmpl, "templates/project/internal/api/http/handler/handler.gotmpl")
	if err != nil {
		panic(err)
	}

	handlerDirPath := fmt.Sprintf("%s/%s/http/handler", rootDir, cfg.OutputPath)
	handlerPath := fmt.Sprintf("%s/%s", handlerDirPath, "/handler.go")
	if _, err = os.Stat(handlerPath); cfg.OpenApiPath != "" && err != nil {
		cmd.logger.Debug("empty handler, add " + specPath)
		buf := bytes.NewBuffer(nil)
		if err = handlerTemplate.Execute(buf, mainSchema); err != nil {
			panic(err)
		}

		b, err := format.Source(buf.Bytes())
		if err != nil {
			cmd.logger.Error("format.Source(buf.Bytes())", zap.Error(err))

			return err
		}

		// Create the directories if they don't exist
		err = os.MkdirAll(handlerDirPath, os.ModePerm)
		if err != nil {
			panic(err)
		}

		err = os.WriteFile(handlerPath, b, 0744)
		if err != nil {
			fmt.Println(err)
			return err
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

	cmd.logger.Debug("output path: " + output)

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
