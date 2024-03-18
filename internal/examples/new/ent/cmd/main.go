package main

import (
	"log"
	"reflect"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

func main() {
	genPackage := "github.com/wheissd/mkgo/internal/examples/new/ent/gen"

	cfg := &gen.Config{
		Features: []gen.Feature{
			//gen.FeatureIntercept,
			//gen.FeatureEntQL,
			//gen.FeatureUpsert,
			//gen.FeaturePrivacy,
			//gen.FeatureSnapshot,
			//gen.FeatureSchemaConfig,
			//gen.FeatureModifier,
		},
		Target:  "./ent/gen",
		Package: genPackage,
		IDType: &field.TypeInfo{Type: field.TypeUUID, RType: &field.RType{
			Kind: reflect.TypeOf(uuid.UUID{}).Kind(),
		}},
	}
	err := entc.TemplateDir("./ent/schema/template")(cfg)
	if err != nil {
		panic(err)
	}

	if err := entc.Generate("./ent/schema", cfg); err != nil {
		log.Fatal("running ent codegen:", err)
	}
}
