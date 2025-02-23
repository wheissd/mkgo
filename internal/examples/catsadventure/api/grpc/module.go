package grpc

import (
	"net"

	"github.com/wheissd/mkgo/internal/examples/catsadventure/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var Module = fx.Module(
	"grpc",
	fx.Provide(
		NewBreedServiceImpl,
		NewCatServiceImpl,
		NewFatherCatServiceImpl,
		NewKittenServiceImpl,
	),
	fx.Invoke(
		func(
			breedImpl *BreedServiceImpl,
			catImpl *CatServiceImpl,
			fathercatImpl *FatherCatServiceImpl,
			kittenImpl *KittenServiceImpl,
			logger *slog.Logger,
			grpcCfg config.GRPC,
		) {
			srv := grpc.NewServer()
			RegisterBreedServiceServer(srv, breedImpl)
			RegisterCatServiceServer(srv, catImpl)
			RegisterFatherCatServiceServer(srv, fathercatImpl)
			RegisterKittenServiceServer(srv, kittenImpl)
			listener, err := net.Listen("tcp", grpcCfg.HTTP.Addr+":"+grpcCfg.HTTP.Port)
			if err != nil {
				logger.Fatal("failed to listen", slog.Any("error", err))
			}
			err = srv.Serve(listener)
		},
	),
)
