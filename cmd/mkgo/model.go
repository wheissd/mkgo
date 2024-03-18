package main

import (
	"bytes"
	"embed"
	"fmt"
	"go/format"
	"os"
	"strings"
	"text/template"

	"github.com/urfave/cli/v2"
)

var (
	//go:embed templates/model.gotmpl
	modelTmpl embed.FS
)

func (cmd *cmd) model(ctx *cli.Context) error {
	type modelConfig struct {
		Name string
	}
	modelName := ctx.Args().First()
	t, err := template.New("model.gotmpl").
		ParseFS(modelTmpl, "templates/model.gotmpl")
	if err != nil {
		return err
	}

	outputFile := fmt.Sprintf("internal/ent/schema/%s.go", strings.ToLower(modelName))

	wd, _ := os.Getwd()
	wdSplitted := strings.Split(wd, "/")
	wdLen := len(wdSplitted)
	for i := wdLen - 1; i > wdLen-5; i-- {
		if wdSplitted[i] == "internal" {
			outputFile = strings.TrimPrefix(outputFile, "internal/")
		}
	}

	cfg := modelConfig{
		Name: modelName,
	}
	buf := bytes.NewBuffer(nil)
	err = t.Execute(buf, cfg)
	if err != nil {
		panic(err)
	}
	var b = buf.Bytes()
	b, err = format.Source(buf.Bytes())
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(outputFile, b, 0744)
	if err != nil {
		panic(err)
	}
	return nil
}
