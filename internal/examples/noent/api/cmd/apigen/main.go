package main

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/wheissd/mkgo/gen"
	"github.com/wheissd/mkgo/options"
)

//go:generate go run github.com/wheissd/gomk/cmd/pre -o ./pre_gen.go -schema_path internal/examples/noent/internal/ent/schema -gen_path internal/examples/noent/ent/gen

func main() {
	gen.Gen(entities, options.OpenapiSchema(func(t *openapi3.T) error { return nil }))
}
