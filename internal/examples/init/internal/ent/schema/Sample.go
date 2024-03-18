package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
)

// Sample holds the schema definition for the Sample entity.
type Sample struct {
	ent.Schema
}

func (Sample) Mixin() []ent.Mixin {
	return []ent.Mixin{}
}

// Fields of the Sample.
func (Sample) Fields() []ent.Field {
	return []ent.Field{}
}

// Edges of the Sample.
func (Sample) Edges() []ent.Edge {
	return []ent.Edge{}
}

// Annotations of the Sample.
func (Sample) Annotations() []schema.Annotation {
	return []schema.Annotation{}
}
