package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/wheissd/mkgo/annotations"
)

// DefaultModel holds the schema definition for the DefaultModel entity.
type DefaultModel struct {
	ent.Schema
}

func (DefaultModel) Mixin() []ent.Mixin {
	return []ent.Mixin{}
}

// Fields of the DefaultModel.
func (DefaultModel) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}),
		field.String("name"),
	}
}

// Edges of the DefaultModel.
func (DefaultModel) Edges() []ent.Edge {
	return []ent.Edge{}
}

// Annotations of the DefaultModel.
func (DefaultModel) Annotations() []schema.Annotation {
	return []schema.Annotation{
		annotations.Entity().
			EnableReadMany(annotations.Modes{"rest_client"}).
			EnableReadOne(annotations.Modes{"rest_client"}),
	}
}
