package main

import (
	"bytes"
	"embed"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"html/template"
	"os"
	"strings"

	"go.uber.org/zap"
)

var (
	//go:embed templates/config.gotmpl
	configTmpl embed.FS
)

type constructor struct {
	Name string
	Type string
}

type genData struct {
	Constructors []constructor
	Imports      []string
	Path         string
}

func (cmd *cmd) genConfig(target string) {
	configTemplate, err := template.ParseFS(configTmpl, "templates/config.gotmpl")
	if err != nil {
		panic(err)
	}

	prsr, err := parser.ParseDir(
		token.NewFileSet(),
		target,
		nil,
		0,
	)
	if err != nil {
		panic(err)
	}

	gData := genData{
		Constructors: make([]constructor, 0),
		Path:         target,
	}

	for _, pkg := range prsr {
		for fName, file := range pkg.Files {
			fName = strings.TrimPrefix(fName, target+"/")
			if fName != "config.go" {
				continue
			}
			for _, decl := range file.Decls {
				if gDecl, ok := decl.(*ast.GenDecl); ok {
					cmd.logger.Debug("decl", zap.String("decl", gDecl.Tok.String()))
					switch gDecl.Tok {
					case token.IMPORT:
						cmd.logger.Debug("decl", zap.Any("specs", gDecl.Specs))
						for _, spec := range gDecl.Specs {
							if importSpec, ok := spec.(*ast.ImportSpec); ok {
								cmd.logger.Debug("decl", zap.Any("importSpec", spec))
								gData.Imports = append(gData.Imports, strings.Trim(importSpec.Path.Value, "\""))
							}
						}
					case token.TYPE:
						for _, spec := range gDecl.Specs {
							if tSpec, ok := spec.(*ast.TypeSpec); ok {
								if tSpec.Name.Name == "Config" {
									typeFields := tSpec.Type.(*ast.StructType).Fields
									for _, field := range typeFields.List {
										fName := field.Names[0].Name
										var (
											fType    *ast.Ident
											selector *ast.SelectorExpr
											ok       bool
										)

										// конфиг лежит в другом пакете
										if selector, ok = field.Type.(*ast.SelectorExpr); ok {
											cfgPkg := selector.X.(*ast.Ident).Name
											gData.Constructors = append(gData.Constructors, constructor{
												Name: fName,
												Type: cfgPkg + "." + selector.Sel.Name,
											})
											continue
										}

										// конфиг лежит в этом же пакете
										if fType, ok = field.Type.(*ast.Ident); !ok {
											continue
										}
										if fType.Obj == nil {
											continue
										}
										gData.Constructors = append(gData.Constructors, constructor{
											Name: fName,
											Type: fType.Obj.Name,
										})
									}
								}
							}
						}
					}
				}
			}
		}
	}

	buf := bytes.NewBuffer(nil)

	cmd.logger.Debug("imports", zap.Any("config imports", gData.Imports))
	err = configTemplate.Execute(buf, gData)
	if err != nil {
		panic(err)
	}

	b, err := format.Source(buf.Bytes())
	if err != nil {
		panic(err)
	}

	outputPath := gData.Path + "/gen.go"
	err = os.WriteFile(outputPath, b, 0744)
	if err != nil {
		panic(err)
	}
	cmd.runCmd("goimports -l -w " + outputPath)
}
