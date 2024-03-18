package main

import (
	"github.com/wheissd/mkgo/gen"
)

//go:generate go run github.com/wheissd/gomk/cmd/pre -o ./pre_gen.go -schema_path internal/examples/new2/ent/schema -gen_path internal/examples/new2/ent/gen

func main() {
	gen.Gen(entities)
}
