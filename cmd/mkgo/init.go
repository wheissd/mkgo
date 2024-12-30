package main

import (
	"errors"
	"log"
	"os"
	"strings"

	"github.com/samber/lo"
	"github.com/urfave/cli/v2"
	"github.com/wheissd/mkgo/internal/parse"
	"go.uber.org/zap"
)

func (cmd *cmd) initProject(ctx *cli.Context) error {
	path := ctx.Args().First()
	if path == "" {
		log.Fatal("please, specify target dir")
	}
	if !strings.HasSuffix(path, "/") {
		path = path + "/"
	}

	targetModuleName := ctx.String("module")
	if targetModuleName == "" {
		splittedPath := strings.Split(strings.TrimSuffix(path, "/"), "/")
		targetModuleName = splittedPath[len(splittedPath)-1]
	}
	mod := parse.GetMod(path)
	cmd.logger.Debug(
		"start mkgo init",
		zap.String("Path", path),
		zap.String("mod", mod),
	)
	pathContainsInternal := strings.Contains(path, "internal")
	if pathContainsInternal {
		for i := range basicDirs {
			basicDirs[i] = strings.TrimPrefix(basicDirs[i], "internal/")
		}
		basicDirs = lo.Filter(basicDirs, func(item string, _ int) bool { return item != "internal" })
	}
	if err := checkDirs(path); err != nil {
		return err
	}
	if err := writeDirs(path); err != nil {
		return err
	}
	toolBasePath := "github.com/wheissd/mkgo"
	schemaPath := path
	if schemaPath == "./" {
		schemaPath = ""
	}
	sch := Schema{
		Tool:         Tool{Module: toolBasePath},
		Module:       mod,
		Path:         schemaPath,
		ProjectPath:  path,
		InternalMode: pathContainsInternal,
		TargetName:   targetModuleName,
	}
	err := writeGoFiles(sch)
	if err != nil {
		return err
	}

	cmd.runVCmd("go get entgo.io/ent")

	cmd.runVCmd("go install github.com/ogen-go/ogen/cmd/ogen@latest")

	return cmd.runVCmd("go get entgo.io/ent/entc")
}

var basicDirs = []string{
	"cmd",
	"cmd/restclient",
	"cmd/grpcadmin",
	"internal",
	"internal/app",
	"internal/db",
	"internal/rest_client",
	"internal/grpc_admin",
	"internal/rest_client/ogen",
	"internal/rest_client/grpc",
	"internal/rest_client/service",
	"internal/rest_client/cmd",
	"internal/rest_client/cmd/apigen",
	"internal/rest_client/http",
	"internal/rest_client/http/handler",
	"internal/grpc_admin/ogen",
	"internal/grpc_admin/grpc",
	"internal/grpc_admin/service",
	"internal/grpc_admin/cmd",
	"internal/grpc_admin/cmd/apigen",
	"internal/grpc_admin/http",
	"internal/grpc_admin/http/handler",
	"internal/ent",
	"internal/ent/schema",
	"internal/ent/schema/template",
	"internal/ent/migrate",
	"internal/ent/migrate/migrations",
	"internal/ent/cmd",
	"internal/config",
	"openapi",
	"proto",
	"mkgo",
}

var ErrTargetNotEmpty = errors.New("target directory is not empty")

func checkDirs(path string) error {
	dirs, err := os.ReadDir(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	for _, dirEntry := range dirs {
		if lo.Contains(basicDirs, dirEntry.Name()) {
			return ErrTargetNotEmpty
		}
	}
	return nil
}

func writeDirs(path string) error {
	if path != "./" {
		if err := os.Mkdir(path, 0744); err != nil && err != os.ErrExist {
			return err
		}
	}
	for _, d := range basicDirs {
		if err := os.Mkdir(path+d, 0744); err != nil {
			return err
		}
	}
	return nil
}
