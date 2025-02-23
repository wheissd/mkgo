package gen

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"reflect"
	"strings"
	"text/template"

	"github.com/samber/lo"
	entadapter "github.com/wheissd/mkgo/internal/adapters/ent"
	"github.com/wheissd/mkgo/internal/cases"
	"github.com/wheissd/mkgo/internal/config"
	"github.com/wheissd/mkgo/internal/entity"
	logger2 "github.com/wheissd/mkgo/internal/logger"
	"github.com/wheissd/mkgo/internal/parse"
	"github.com/wheissd/mkgo/lib"
	genoptions "github.com/wheissd/mkgo/options"
)

var fns = template.FuncMap{
	"last": func(x int, a interface{}) bool {
		return x == reflect.ValueOf(a).Len()-1
	},
	"notLast": func(x int, a interface{}) bool {
		return x != reflect.ValueOf(a).Len()-1
	},
	"hasFormat":                 hasFormat,
	"format":                    paramFormat,
	"pascal":                    cases.Pascal,
	"camel":                     cases.Camel,
	"snake":                     cases.Snake,
	"lower":                     strings.ToLower,
	"setReqFormat":              setReqFormat,
	"setEdgeReqFormat":          setEdgeReqFormat,
	"updateReqFormat":           updateReqFormat,
	"isEnum":                    isEnum,
	"needReadOneOp":             needReadOneOp,
	"needReadManyOp":            needReadManyOp,
	"needCreateOp":              needCreateOp,
	"needUpdateOp":              needUpdateOp,
	"needDeleteOp":              needDeleteOp,
	"hasOps":                    hasOps,
	"isFieldPublic":             isFieldPublic,
	"isIDField":                 isIDField,
	"fieldType":                 fieldType,
	"fieldDefault":              fieldDefault,
	"fieldTypeIs":               fieldTypeIs,
	"updateFieldType":           updateFieldType,
	"sprintf":                   fmt.Sprintf,
	"dict":                      dict,
	"concat":                    concat,
	"techField":                 techField,
	"protoToServiceField":       protoToServiceField,
	"protoToServiceFieldUpdate": protoToServiceFieldUpdate,
	"protoToServiceFieldFilter": protoToServiceFieldFilter,
	"serviceToProtoField":       serviceToProtoField,
	"add":                       func(v, v2 int) int { return v + v2 },
	"maxInt": func(v1, v2 int) int {
		if v1 > v2 {
			return v1
		}
		return v2
	},
	"needFilter": needFilter,
	"needEntity": needEntity,
}

const (
	TransportHTTP    = "http"
	TransportDefault = ""
	TransportGRPC    = "grpc"
)

func Gen(entities []lib.PreEntity, options ...genoptions.GenOption) {
	entPath := flag.String("ent_path", "internal/ent/gen", "ent gen directory path")
	cfgPath := flag.String("cfg_path", "mkgo_config.yaml", "api gen config path")
	cfgRaw := flag.String("cfg", "{}", "api gen config in json format")
	cfgMode := flag.String("mode", "default", "api gen mode, specified in mkgo_config")
	flag.Parse()

	logger := logger2.Get(slog.LevelError)
	defer func() {
		err := recover()
		if err != nil {
			if e, ok := err.(error); ok {
				logger.Error(
					"recover()",
					slog.Any("error", e),
					slog.Bool("isErr", true),
				)
			} else {
				logger.Error("recover()", slog.Any("err", err), slog.Bool("isErr", false))
			}
			return
		}
	}()
	var cfg config.GenConfig
	if *cfgRaw != "{}" {
		err := json.Unmarshal([]byte(*cfgRaw), &cfg)
		if err != nil {
			panic(err)
		}
	} else {
		cfg.Parse(*cfgPath)
	}
	logger.Info("start apigen", slog.Any("cfg", cfg))
	*entPath = strings.TrimPrefix(*entPath, "/")

	for _, apiCfg := range cfg.APIs {
		if apiCfg.Mode != *cfgMode {
			continue
		}
		sch := makeSchema(logger, apiCfg, entPath)

		entadapter.Parse(entities, sch, apiCfg.Mode)
		logger.Debug("entities:")
		entityNames := lo.Map(sch.Entities, func(v *entity.Entity, _ int) string {
			return v.Name
		})
		logger.Debug(strings.Join(entityNames, " "))

		err := updateFieldsOrder(sch)
		if err != nil {
			panic(err)
		}

		switch apiCfg.Transport {
		case TransportHTTP, TransportDefault:
			oapi(logger, sch, options...)
		case TransportGRPC:
			proto(logger, apiCfg, sch)
		}

		logger.Debug("start write go files")
		writeGoFiles(logger, sch)
	}
}

func makeSchema(logger *slog.Logger, cfg config.GenConfigItem, entPath *string) *entity.Schema {
	wd, _ := os.Getwd()
	pkg, lvl := parse.GetModLevel(wd, 0)
	wdSplitted := strings.Split(wd, "/")
	rootPkg := strings.Join(lo.Filter(wdSplitted, func(item string, i int) bool {
		return i > len(wdSplitted)-lvl-1
	}), "/")
	if rootPkg == "" {
		rootPkg = "internal"
	}
	if !strings.HasSuffix(rootPkg, "/") {
		rootPkg = rootPkg + "/"
	}
	logger.Debug(
		"makeSchema",
		slog.String("wd", wd),
		slog.String("entPath", *entPath),
		slog.String("pkg", pkg),
		slog.String("rootPkg", rootPkg),
	)
	sch := &entity.Schema{
		Cfg:          cfg,
		Pkg:          pkg,
		EntPath:      *entPath,
		WD:           wd,
		RootPkg:      rootPkg,
		ProtoImports: make([]string, 0),
	}
	return sch
}

func oapi(logger *slog.Logger, sch *entity.Schema, options ...genoptions.GenOption) {
	logger.Debug("gen openapi")
	buf := bytes.NewBuffer(nil)
	externalSchema := []genoptions.OpenapiSchemaOption{}
	for i := range options {
		if externalSch, ok := options[i].(genoptions.OpenapiSchemaOption); ok {
			externalSchema = append(externalSchema, externalSch)
		}
	}
	oapi := genOpenapi(logger, sch, sch.Cfg, externalSchema...)
	oapiJSON, err := json.Marshal(oapi)
	if err != nil {
		panic(err)
	}
	indentedOapi := bytes.NewBuffer(nil)
	_ = os.Remove(sch.Cfg.OpenApiPath)
	err = json.Indent(indentedOapi, oapiJSON, "", "  ")
	if err != nil {
		indentErr := err
		err = os.WriteFile(sch.Cfg.OpenApiPath, buf.Bytes(), 0744)
		if err != nil {
			panic(err)
		}
		panic(indentErr)
	}
	err = os.WriteFile(sch.Cfg.OpenApiPath, indentedOapi.Bytes(), 0744)
	if err != nil {
		panic(err)
	}
}
