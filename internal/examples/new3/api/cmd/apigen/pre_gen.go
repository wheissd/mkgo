package main

import (
	"github.com/wheissd/mkgo/internal/examples/new3/ent/gen"
	. "github.com/wheissd/mkgo/internal/examples/new3/ent/schema"
	"github.com/wheissd/mkgo/lib"
)

var entities = []lib.PreEntity{
	{
		Schema: Human{},
		Model:  gen.Human{},
	},
}
