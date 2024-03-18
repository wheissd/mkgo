package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Human holds the schema definition for the Human entity.
type Human struct {
	ent.Schema
}

func (Human) Mixin() []ent.Mixin {
	return []ent.Mixin{}
}

// Fields of the Human.
func (Human) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}),
		field.String("name"),
	}
}

// Edges of the Human.
func (Human) Edges() []ent.Edge {
	return []ent.Edge{}
}

// Annotations of the Human.
func (Human) Annotations() []schema.Annotation {
	return []schema.Annotation{}
}
