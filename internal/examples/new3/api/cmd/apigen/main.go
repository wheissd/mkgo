package main

import (
	"github.com/wheissd/mkgo/gen"
)

//go:generate go run github.com/wheissd/gomk/cmd/pre -o . -schema_path internal/examples/new3/ent/schema -gen_path internal/examples/new3/ent/gen

func main() {
	gen.Gen(entities)
}
