package options

import (
	"github.com/getkin/kin-openapi/openapi3"
)

type GenOption interface {
	GenOption()
}

type OpenapiSchemaOption struct {
	SchemaCallback func(t *openapi3.T) error
}

func OpenapiSchema(f func(t *openapi3.T) error) OpenapiSchemaOption {
	return OpenapiSchemaOption{SchemaCallback: f}
}

func (o OpenapiSchemaOption) GenOption() {}
