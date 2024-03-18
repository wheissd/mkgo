package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const toolName = "mkgo"

func main() {
	cmd := cmd{}
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
		cmd.logger.Error("run failed", zap.Error(err))
	}
}

func (cmd *cmd) initLogger(ctx *cli.Context) error {
	logLevel := zap.ErrorLevel
	if ctx.IsSet("d") {
		logLevel = zap.DebugLevel
	}
	pe := zap.NewDevelopmentEncoderConfig()
	pe.EncodeLevel = zapcore.CapitalColorLevelEncoder
	pe.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(pe),
		zapcore.AddSync(os.Stdout),
		logLevel,
	)

	cmd.noTraceLogger = zap.New(core, zap.AddCaller())
	cmd.logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
	return nil
}

func (cmd *cmd) logErr(ctx *cli.Context) error {
	err := ctx.Err()
	if err != nil {
		cmd.logger.Error("run err", zap.Error(err))
	}
	return nil
}
