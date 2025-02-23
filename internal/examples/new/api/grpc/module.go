package grpc

import (
	"net"

	"github.com/wheissd/mkgo/internal/examples/new/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var Module = fx.Module(
	"grpc",
	fx.Provide(
		New,
	),
	fx.Invoke(
		func(
			impl *DefaultModelServiceImpl,
			logger *slog.Logger,
			grpcCfg config.GRPC,
		) {
			srv := grpc.NewServer()
			RegisterDefaultModelServiceServer(srv, impl)
			listener, err := net.Listen("tcp", grpcCfg.HTTP.Addr+":"+grpcCfg.HTTP.Port)
			if err != nil {
				logger.Fatal("failed to listen", slog.Any("error", err))
			}
			err = srv.Serve(listener)
		},
	),
)
