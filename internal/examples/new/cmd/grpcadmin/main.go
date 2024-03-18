package main

import (
	"github.com/wheissd/mkgo/internal/examples/new/api/grpc"
	"github.com/wheissd/mkgo/internal/examples/new/app"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		append(
			app.DI(),
			grpc.Module,
		)...,
	)

	<-app.Wait()
}
