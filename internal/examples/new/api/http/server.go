package http

import (
	"context"
	"log"
	"net/http"

	"github.com/wheissd/mkgo/internal/examples/new/api/http/handler"
	"github.com/wheissd/mkgo/internal/examples/new/api/ogen"
	"github.com/wheissd/mkgo/internal/examples/new/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func start(
	cfg config.Config,
	srv *ogen.Server,
	logger *slog.Logger,
	shutdowner fx.Shutdowner,
) {
	go func() {
		err := http.ListenAndServe(
			cfg.API.HTTP.Addr+":"+cfg.API.HTTP.Port,
			srv,
		)
		if err != nil {
			logger.Error("unable to start admin server", slog.Any("error", err))
			err := shutdowner.Shutdown()
			if err != nil {
				panic(err)
			}
		}
	}()
}

func srv(h *handler.Handler, logger *slog.Logger) *ogen.Server {
	// start listening.
	srv, err := ogen.NewServer(
		h,
		ogen.WithMethodNotAllowed(func(w http.ResponseWriter, r *http.Request, allowed string) {
			status := http.StatusMethodNotAllowed
			if r.Method == "OPTIONS" {
				status = http.StatusNoContent
			} else {
				w.Header().Set("Allow", allowed)
			}
			w.WriteHeader(status)
			logger.Debug("Method not allowed",
				//slog.Int("http.status", resp.Type),
				slog.String("uri", r.RequestURI),
				slog.Any("method", r.Method),
			)
		}),
		ogen.WithErrorHandler(func(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {
			logger.Debug("Ogen error",
				slog.String("uri", r.RequestURI),
				slog.Any("method", r.Method),
				slog.Any("error", err),
			)
		}),
	)
	if err != nil {
		log.Fatal(err)
	}

	return srv
}

var Module = fx.Module(
	"http",
	fx.Provide(
		srv,
		handler.New,
	),
	fx.Invoke(start),
)
