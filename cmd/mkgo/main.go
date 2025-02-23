package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/urfave/cli/v2"
	"github.com/wheissd/mkgo/internal/logger"

	"log/slog"
)

const toolName = "mkgo"

func main() {
	cmd := cmd{
		depCheckClient: &http.Client{Timeout: 5 * time.Second},
	}
	cmds := []*cli.Command{
		{
			Name:   "init",
			Usage:  "init project",
			Action: cmd.initProject,
		},
		{
			Name:   "model",
			Usage:  "make model",
			Action: cmd.model,
		},
		{
			Name:   "generate",
			Usage:  "update project files",
			Action: cmd.generate,
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:  "f",
					Usage: "force regenerate all",
				},
			},
		},
		{
			Name:   "version",
			Usage:  "get tool version",
			Action: cmd.version,
		},
	}
	for i := range cmds {
		cmds[i].Before = cmd.initLogger
		cmds[i].After = cmd.logErr
		cmds[i].Flags = append(cmds[i].Flags, &cli.BoolFlag{
			Name:  "d",
			Usage: "debug mode",
		})
	}

	err := (&cli.App{
		Name:  "tool",
		Usage: fmt.Sprintf("%s is a framework to generate web services", toolName),
		Action: func(ctx *cli.Context) error {
			ctx.String("wd")
			return nil
		},
		Commands: cmds}).Run(os.Args)
	if err != nil {
		cmd.logger.Error("run failed", slog.Any("error", err))
	}
}

func (cmd *cmd) initLogger(ctx *cli.Context) error {
	logLevel := slog.LevelError
	if ctx.IsSet("d") {
		logLevel = slog.LevelDebug
	}

	//pe := slog.NewDevelopmentEncoderConfig()
	//pe.EncodeLevel = zapcore.CapitalColorLevelEncoder
	//pe.EncodeTime = zapcore.ISO8601TimeEncoder
	//core := zapcore.NewCore(
	//	zapcore.NewConsoleEncoder(pe),
	//	zapcore.AddSync(os.Stdout),
	//	logLevel,
	//)

	//cmd.noTraceLogger = slog.New(core, slog.AddCaller())
	cmd.logger = logger.Get(logLevel)
	return nil
}

func (cmd *cmd) logErr(ctx *cli.Context) error {
	err := ctx.Err()
	if err != nil {
		cmd.logger.Error("run err", slog.Any("error", err))
	}
	return nil
}
