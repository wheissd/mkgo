package main

import (
	"github.com/wheissd/mkgo/internal/examples/catsadventure/ent/gen"
	. "github.com/wheissd/mkgo/internal/examples/catsadventure/ent/schema"
	"github.com/wheissd/mkgo/lib"
)

var entities = []lib.PreEntity{
	{
		Schema: Breed{},
		Model:  gen.Breed{},
	},
	{
		Schema: Cat{},
		Model:  gen.Cat{},
	},
	{
		Schema: FatherCat{},
		Model:  gen.FatherCat{},
	},
	{
		Schema: Kitten{},
		Model:  gen.Kitten{},
	},
}
