package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/wheissd/mkgo/lib/ent/mixin"
)

// Breed holds the schema definition for the Breed entity.
type Breed struct {
	ent.Schema
}

func (Breed) Mixin() []ent.Mixin {
	return append([]ent.Mixin{}, mixin.Default()...)
}

// Fields of the Breed.
func (Breed) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

// Edges of the Breed.
func (Breed) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("cats", Cat.Type),
	}
}
