package main

import (
	"github.com/wheissd/mkgo/internal/examples/new2/ent/gen"
	. "github.com/wheissd/mkgo/internal/examples/new2/ent/schema"
	"github.com/wheissd/mkgo/lib"
)

var entities = []lib.PreEntity{
	{
		Schema: Human{},
		Model:  gen.Human{},
	},
}
