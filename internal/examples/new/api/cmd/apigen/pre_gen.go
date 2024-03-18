package main

import (
	"github.com/wheissd/mkgo/internal/examples/new/ent/gen"
	. "github.com/wheissd/mkgo/internal/examples/new/ent/schema"
	"github.com/wheissd/mkgo/lib"
)

var entities = []lib.PreEntity{
	{
		Schema: DefaultModel{},
		Model:  gen.DefaultModel{},
	},
}
