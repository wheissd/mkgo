package entity

import (
	"github.com/samber/lo"
	"github.com/wheissd/mkgo/annotations"
	"github.com/wheissd/mkgo/lib"
)

type Parser interface {
	Parse([]lib.PreEntity, *Schema)
}

type Entity struct {
	Model          any
	Path           string
	Name           string
	NeedFlat       bool
	Fields         []Field
	Edges          []*Edge
	HasReadEdges   bool
	HasCreateEdges bool
	HasUpdateEdges bool
	Config         *annotations.EntityConfig
}

func (e Entity) GetIDField() Field {
	return lo.Filter(e.Fields, func(item Field, _ int) bool {
		return item.Name == "id"
	})[0]
}

type Openapi struct {
	Type      string
	Format    string
	HasFormat bool
}
