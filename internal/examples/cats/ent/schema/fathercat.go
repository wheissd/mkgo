package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/wheissd/mkgo/lib/ent/mixin"
)

// FatherCat holds the schema definition for the FatherCat entity.
type FatherCat struct {
	ent.Schema
}

func (FatherCat) Mixin() []ent.Mixin {
	return append([]ent.Mixin{}, mixin.Default()...)
}

// Fields of the FatherCat.
func (FatherCat) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

// Edges of the FatherCat.
func (FatherCat) Edges() []ent.Edge {
	return nil
}
