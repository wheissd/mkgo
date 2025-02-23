package main

import (
	"bytes"
	"embed"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"html/template"
	"log/slog"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

var (
	//go:embed templates/config.gotmpl
	configTmpl embed.FS
)

type constructor struct {
	Name string
	Type string
}

type importData struct {
	Path string
	As   *string
}

type genData struct {
	Constructors []constructor
	Imports      []importData
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
					cmd.logger.Debug("decl", slog.String("decl", gDecl.Tok.String()))
					switch gDecl.Tok {
					case token.IMPORT:
						cmd.logger.Debug("decl", slog.Any("specs", gDecl.Specs))
						for _, spec := range gDecl.Specs {
							if importSpec, ok := spec.(*ast.ImportSpec); ok {
								cmd.logger.Debug("decl", slog.Any("importSpec", spec))
								imp := importSpec.Path.Value
								imD := importData{Path: strings.Trim(imp, "\"")}
								if importSpec.Name != nil {
									imD.As = lo.ToPtr(importSpec.Name.String())
								}
								gData.Imports = append(gData.Imports, imD)
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

	cmd.logger.Debug("imports", slog.Any("config imports", gData.Imports))
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
	cmd.runVCmd("goimports -l -w " + outputPath)
}

//func (cmd *cmd) getCorrectImportData(gData genData, imD importData) importData {
//	if _, found := lo.Find(gData.Imports, func(i importData) bool {
//		foundPath := i.Path
//		if i.As != nil {
//			foundPath = *i.As
//		}
//		foundSplitted := strings.Split(foundPath, "/")
//		foundPath = foundSplitted[len(foundSplitted)-1]
//		imDPath := imD.Path
//		if imD.As != nil {
//			imDPath = *imD.As
//		}
//		imDPathSplitted := strings.Split(imDPath, "/")
//		s := imDPathSplitted[len(imDPathSplitted)-1]
//		cmd.logger.Debug("getCorrectImportData", slog.String("s", s), slog.String("foundPath", foundPath))
//		return s == foundPath
//	}); found {
//		s := imD.Path
//		if imD.As != nil {
//			s = *imD.As
//		}
//		next := generateNextNumericSuffix(s)
//		return cmd.getCorrectImportData(gData, importData{Path: imD.Path, As: &next})
//	}
//	return importData{
//		Path: imD.Path,
//		As:   imD.As,
//	}
//}

var strNumericTailRegexp = regexp.MustCompile(`(.*?)(\d+)$`)

func generateNextNumericSuffix(input string) string {
	parts := strings.Split(input, "/")
	lastPart := parts[len(parts)-1]
	matches := strNumericTailRegexp.FindStringSubmatch(lastPart)

	if len(matches) < 3 {
		// No numeric tail found, return input with "1" appended
		return lastPart + "1"
	}

	// Extract the prefix and numeric tail
	prefix := matches[1]
	numberStr := matches[2]

	// Parse the numeric tail
	number, _ := strconv.Atoi(numberStr)

	// Increment the number
	number++

	// Format the new string with the same number of digits
	newNumberStr := fmt.Sprintf("%0*d", len(numberStr), number)

	// Return the new string
	return prefix + newNumberStr
}
