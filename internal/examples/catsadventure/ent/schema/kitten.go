package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/wheissd/mkgo/lib/ent/mixin"
)

// Kitten holds the schema definition for the Kitten entity.
type Kitten struct {
	ent.Schema
}

func (Kitten) Mixin() []ent.Mixin {
	return append([]ent.Mixin{}, mixin.Default()...)
}

// Fields of the Kitten.
func (Kitten) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.UUID("mother_id", uuid.UUID{}),
	}
}

// Edges of the Kitten.
func (Kitten) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("mother", Cat.Type).
			Required().
			Unique().
			Field("mother_id").
			Ref("kittens"),
	}
}
