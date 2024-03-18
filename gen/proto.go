package gen

import (
	"bytes"
	"embed"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/wheissd/mkgo/internal/config"
	"github.com/wheissd/mkgo/internal/entity"
	"go.uber.org/zap"
)

var (
	//go:embed templates/proto/*tmpl
	protoTmpl embed.FS
)

func proto(logger *zap.Logger, cfg config.GenConfigItem, sch *entity.Schema) {
	logger.Debug("gen proto")
	fPath := "templates/proto/entities.prototmpl"
	fName := "entities.prototmpl"
	var err error
	protoTemplate, err := template.New(fName).Funcs(fns).
		ParseFS(protoTmpl, fPath)
	if err != nil {
		panic(err)
	}

	err = cleanProtoDir(cfg)
	if err != nil {
		panic(err)
	}

	buf := bytes.NewBuffer(nil)

	err = protoTemplate.Execute(buf, sch)
	if err != nil {
		panic(err)
	}
	var b = buf.Bytes()
	outputFile := sch.Cfg.ProtoPath + "/entities_gen.proto"

	logger.Debug(fmt.Sprintf("write proto file: %s\n", outputFile))

	err = os.WriteFile(outputFile, b, 0744)
	if err != nil {
		panic(err)
	}

}

func cleanProtoDir(cfg config.GenConfigItem) error {
	path := "./proto"
	dir, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for _, f := range dir {
		if strings.HasSuffix(f.Name(), "proto") {
			err = os.Remove(path + "/" + f.Name())
			if err != nil {
				return err
			}
		}
	}
	return nil
}
