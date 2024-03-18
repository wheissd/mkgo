package app

import (
	"context"

	"github.com/wheissd/mkgo/internal/examples/new/api/service"
	"github.com/wheissd/mkgo/internal/examples/new/config"
	"github.com/wheissd/mkgo/internal/examples/new/db"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func DI() []fx.Option {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	return []fx.Option{
		// app context
		fx.Provide(
			func() context.Context {
				return context.Background()
			},
			func() *zap.Logger {
				return logger
			},
		),
		db.Module,
		service.Module,
		config.Module,
	}
}
