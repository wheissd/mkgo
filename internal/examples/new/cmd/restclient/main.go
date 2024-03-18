package main

import (
	"github.com/wheissd/mkgo/internal/examples/new/api/http"
	"github.com/wheissd/mkgo/internal/examples/new/app"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		append(
			app.DI(),
			http.Module,
		)...,
	)

	<-app.Wait()
}
