package schema

import (
	"ariga.io/atlas/sql/postgres"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/wheissd/mkgo/annotations"
	"github.com/wheissd/mkgo/lib/ent/mixin"
)

// Cat holds the schema definition for the Cat entity.
type Cat struct {
	ent.Schema
}

func (Cat) Mixin() []ent.Mixin {
	return append([]ent.Mixin{}, mixin.Default()...)
}

func (Cat) Annotations() []schema.Annotation {
	return []schema.Annotation{
		annotations.Entity().DisableDelete(annotations.Modes{"api"}),
	}
}

// Fields of the Cat.
func (Cat) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.UUID("breed_id", uuid.UUID{}).
			Comment("Идентификатор локации"),
		field.Int64("speed").Annotations(
			annotations.Field().SetPublic(),
		),
		field.Time("date_from").
			SchemaType(map[string]string{
				dialect.Postgres: postgres.TypeDate,
			}).
			Comment("Дата начала"),
		field.Enum("type").
			Values(
				"merch",
				"hotel",
				"tournament",
			).
			Comment("Тип"),
		field.Enum("other_type").
			Values(
				"merch",
				"hotel",
				"tournament",
			).
			Comment("Другой тип"),
	}
}

// Edges of the Cat.
func (Cat) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("kittens", Kitten.Type).Annotations(
			annotations.Edge().EnableRead(),
		),
		edge.From("breed", Breed.Type).
			Required().
			Unique().
			Field("breed_id").
			Ref("cats"),
	}
}
