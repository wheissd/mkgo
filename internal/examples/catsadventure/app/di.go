package app

import (
	"context"

	"github.com/wheissd/mkgo/internal/examples/catsadventure/api/service"
	"github.com/wheissd/mkgo/internal/examples/catsadventure/config"
	"github.com/wheissd/mkgo/internal/examples/catsadventure/db"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func DI() []fx.Option {
	logger, err := slog.NewDevelopment()
	if err != nil {
		panic(err)
	}

	return []fx.Option{
		// app context
		fx.Provide(
			func() context.Context {
				return context.Background()
			},
			func() *slog.Logger {
				return logger
			},
		),
		db.Module,
		service.Module,
		config.Module,
	}
}
