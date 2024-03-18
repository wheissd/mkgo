package gen

import (
	. "github.com/getkin/kin-openapi/openapi3"
	"github.com/wheissd/mkgo/internal/entity"
	"go.uber.org/zap"
)

func baseProp(f entity.Field) *SchemaRef {
	t := f.Openapi.Type
	if t == "" {
		t = f.Type.Type.OApiTypeName()
	}
	prop := &SchemaRef{
		Value: &Schema{
			Type:   t,
			Format: paramFormat(f),
		},
	}
	if f.Type.Type == entity.TypeEnum {
		prop.Value.Enum = f.Enum
	}
	return prop
}

func readProperties(logger *zap.Logger, sch *entity.Schema, e *entity.Entity) Schemas {
	res := Schemas{}
	for _, f := range e.Fields {
		logger.Debug("fields")
		if isFieldPublic(sch, f) {
			logger.Debug("field is public", zap.String("name", f.Name))
			res[f.Name] = baseProp(f)
		}
	}
	if e.HasReadEdges {
		edges := Schemas{}
		for _, edge := range e.Edges {
			edges[edge.Name] = edgeProp(edge, schemaEntityNameAsIs)
		}
		res["edges"] = &SchemaRef{
			Value: &Schema{
				Type:       TypeObject,
				Properties: edges,
			},
		}
	}
	return res
}

func readPropertiesRequired(e *entity.Entity) []string {
	res := []string{"id"}
	for _, f := range e.Fields {
		if f.Required {
			res = append(res, f.Name)
		}
	}
	if e.HasReadEdges {
		res = append(res, "edges")
	}
	return res
}

func edgeProp(e *entity.Edge, entityNameFunc func(string) string) *SchemaRef {
	res := &SchemaRef{}
	edgeRef := NewSchemaRef("#/components/schemas/"+entityNameFunc(e.EntityName), nil)
	switch e.Type {
	case entity.EdgeO2M, entity.EdgeM2M:
		res.Value = &Schema{
			Type:  TypeArray,
			Items: edgeRef,
		}
	case entity.EdgeO2O, entity.EdgeM2O:
		res = edgeRef
	}
	return res
}

func edgeCreateProp(e *entity.Edge, entityNameFunc func(string) string) *SchemaRef {
	res := &SchemaRef{}
	edgeRef := NewSchemaRef("#/components/schemas/Create"+entityNameFunc(e.EntityName), nil)
	switch e.Type {
	case entity.EdgeO2M, entity.EdgeM2M:
		res.Value = &Schema{
			Type:  TypeArray,
			Items: edgeRef,
		}
	case entity.EdgeO2O, entity.EdgeM2O:
		res = edgeRef
	}
	return res
}

func createProperties(e *entity.Entity) Schemas {
	res := Schemas{}
	for _, f := range e.Fields {
		res[f.Name] = baseProp(f)
	}
	if len(e.Edges) != 0 {
		edges := Schemas{}
		for _, edge := range e.Edges {
			edges[edge.Name] = edgeCreateProp(edge, schemaEntityNameAsIs)
		}
		res["edges"] = &SchemaRef{
			Value: &Schema{
				Type:       TypeObject,
				Properties: edges,
			},
		}
	}
	return res
}

func createPropertiesRequired(e *entity.Entity) []string {
	res := []string{}
	for _, f := range e.Fields {
		if f.Required {
			res = append(res, f.Name)
		}
	}
	if len(e.Edges) != 0 {
		//edges := Schemas{}
		//for _, edge := range e.Edges {
		//
		//}
	}
	return res
}

func updateProperties(e *entity.Entity) Schemas {
	res := Schemas{}
	for _, f := range e.Fields {
		res[f.Name] = baseProp(f)
	}
	updatePropModifier(e, res)
	return res
}

func updatePropModifier(e *entity.Entity, prop Schemas) {
	if len(e.Edges) != 0 {
		for _, edge := range e.Edges {
			if edge.Type == entity.EdgeO2M || edge.Type == entity.EdgeM2M {
				prop["edges_delete"] = &SchemaRef{
					Value: &Schema{Type: TypeString, Format: "uuid"},
				}
			}
		}
	}
}
