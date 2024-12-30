package main

import (
	"embed"
	"fmt"
	"strings"

	"github.com/urfave/cli/v2"
)

//go:embed version
var version embed.FS

func (cmd *cmd) version(ctx *cli.Context) error {
	f, err := version.ReadFile("version")
	if err != nil {
		return err
	}
	fmt.Print(fmt.Sprintf("mkgo version %s", strings.TrimPrefix(string(f), "v")))
	return nil
}
