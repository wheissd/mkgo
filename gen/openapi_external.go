package gen

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/wheissd/mkgo/options"
)

func schemaExternal(root *openapi3.T, externalConfig ...options.OpenapiSchemaOption) {
	for i := range externalConfig {
		err := externalConfig[i].SchemaCallback(root)
		if err != nil {
			panic(err)
		}
	}
}
