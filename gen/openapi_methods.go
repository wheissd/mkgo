package gen

import (
	"fmt"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/samber/lo"
	"github.com/wheissd/mkgo/internal/cases"
	"github.com/wheissd/mkgo/internal/entity"
)

func makeWithParam(edges []*entity.Edge) *openapi3.ParameterRef {
	var (
		hasReadEdges             bool
		accesibleFirstLevelEdges = make([]string, 0)
	)
	for _, edge := range edges {
		if edge.WithRead {
			hasReadEdges = true
			accesibleFirstLevelEdges = append(accesibleFirstLevelEdges, edge.Name)
		}
	}
	if hasReadEdges {
		return &openapi3.ParameterRef{
			Value: &openapi3.Parameter{
				Name: "with",
				In:   openapi3.ParameterInQuery,
				Description: fmt.Sprintf(
					"return result with edges. accesible first level edges: %s;"+
						" nested edges format: entity.edge,entity.edge.second_level_edge",
					strings.Join(accesibleFirstLevelEdges, ", "),
				),
				Required: false,
				Schema: &openapi3.SchemaRef{
					Value: &openapi3.Schema{
						Type: openapi3.TypeString,
					},
				},
			},
		}
	}
	return nil
}

func makeGetList(sch *entity.Schema, entityItem *entity.Entity) *openapi3.Operation {
	params := openapi3.Parameters{}
	for _, f := range entityItem.Fields {
		t := f.Openapi.Type
		if t == "" {
			t = f.Type.Type.OApiTypeName()
		}
		params = append(params, &openapi3.ParameterRef{
			Value: &openapi3.Parameter{
				In:          openapi3.ParameterInQuery,
				Name:        "filter" + cases.Pascal(f.Name),
				Description: "Filter by " + f.Name + " field",
				Schema: &openapi3.SchemaRef{
					Value: &openapi3.Schema{
						Type:   t,
						Format: paramFormat(f),
					},
				},
			},
		})
	}
	for _, edge := range entityItem.Edges {
		params = append(params, &openapi3.ParameterRef{
			Value: &openapi3.Parameter{
				In:          openapi3.ParameterInQuery,
				Name:        cases.Camel("filter_"+edge.Name) + "Id",
				Description: "Filter by " + edge.EntityName + " relation id",
				Schema: &openapi3.SchemaRef{
					Value: &openapi3.Schema{
						Type:   openapi3.TypeString,
						Format: "uuid",
					},
				},
			},
		})
	}
	if withParam := makeWithParam(entityItem.Edges); withParam != nil {
		params = append(params, withParam)
	}

	return &openapi3.Operation{
		Tags: []string{
			entityItem.Name,
		},
		Summary:     fmt.Sprintf("List for %ss", entityItem.Name),
		Description: fmt.Sprintf("List for %ss", entityItem.Name),
		OperationID: fmt.Sprintf("list%s", entityItem.Name),
		Parameters: append(openapi3.Parameters{
			{
				Value: &openapi3.Parameter{
					Name:        "page",
					In:          openapi3.ParameterInQuery,
					Description: "what page to render",
					Schema: &openapi3.SchemaRef{
						Value: &openapi3.Schema{
							Type: openapi3.TypeInteger,
							Min:  lo.ToPtr(float64(1)),
						},
					},
				},
			},
			{
				Value: &openapi3.Parameter{
					Name:        "itemsPerPage",
					In:          openapi3.ParameterInQuery,
					Description: "item count to render per page",
					Schema: &openapi3.SchemaRef{
						Value: &openapi3.Schema{
							Type: openapi3.TypeInteger,
							Min:  lo.ToPtr(float64(1)),
							Max:  lo.ToPtr(float64(255)),
						},
					},
				},
			}}, params...),
		Responses: openapi3.Responses{
			"200": &openapi3.ResponseRef{
				Value: &openapi3.Response{
					Description: lo.ToPtr(fmt.Sprintf("result %s options headers", entityItem.Name)),
					Headers: openapi3.Headers{
						"Content-Range": {
							Value: &openapi3.Header{
								Parameter: openapi3.Parameter{
									Required: true,
									Schema: &openapi3.SchemaRef{
										Value: &openapi3.Schema{
											Type:        openapi3.TypeInteger,
											Description: "Total items count",
										},
									},
								},
							},
						},
					},
					Content: openapi3.Content{
						"application/json": &openapi3.MediaType{
							Schema: &openapi3.SchemaRef{
								Ref: fmt.Sprintf("#/components/schemas/%sList", entityItem.Name),
							},
						},
					},
				},
			},
			"400": &openapi3.ResponseRef{
				Ref: "#/components/responses/400",
			},
			"404": &openapi3.ResponseRef{
				Ref: "#/components/responses/404",
			},
			"409": &openapi3.ResponseRef{
				Ref: "#/components/responses/409",
			},
			"500": &openapi3.ResponseRef{
				Ref: "#/components/responses/500",
			},
			"default": &openapi3.ResponseRef{
				Ref: "#/components/responses/Error",
			},
		},
	}
}

func makeCreate(entity *entity.Entity) *openapi3.Operation {
	return &openapi3.Operation{
		Tags: []string{
			entity.Name,
		},
		Summary:     fmt.Sprintf("Create a new %s", entity.Name),
		Description: fmt.Sprintf("Create a new %s  and persists it to storage.", entity.Name),
		OperationID: fmt.Sprintf("create%s", entity.Name),
		RequestBody: &openapi3.RequestBodyRef{
			Value: &openapi3.RequestBody{
				Description: fmt.Sprintf("%s to create", entity.Name),
				Content: openapi3.Content{
					"application/json": {
						Schema: openapi3.NewSchemaRef(fmt.Sprintf("#/components/schemas/Create%s", entity.Name), nil),
					},
				},
				Required: true,
			},
		},
		Responses: openapi3.Responses{
			"200": &openapi3.ResponseRef{
				Value: &openapi3.Response{
					Description: lo.ToPtr(fmt.Sprintf("%screated", entity.Name)),
					Content: openapi3.Content{
						"application/json": &openapi3.MediaType{
							Schema: &openapi3.SchemaRef{
								Ref: fmt.Sprintf("#/components/schemas/%s", entity.Name),
							},
						},
					},
				},
			},
			"400": &openapi3.ResponseRef{
				Ref: "#/components/responses/400",
			},
			"404": &openapi3.ResponseRef{
				Ref: "#/components/responses/404",
			},
			"409": &openapi3.ResponseRef{
				Ref: "#/components/responses/409",
			},
			"500": &openapi3.ResponseRef{
				Ref: "#/components/responses/500",
			},
			"default": &openapi3.ResponseRef{
				Ref: "#/components/responses/Error",
			},
		},
	}
}

func makeUpdate(entity *entity.Entity) *openapi3.Operation {
	return &openapi3.Operation{
		Tags: []string{
			entity.Name,
		},
		Summary:     fmt.Sprintf("Update %s", entity.Name),
		Description: fmt.Sprintf("Update %s  and persists it to storage.", entity.Name),
		OperationID: fmt.Sprintf("update%s", entity.Name),
		Parameters: openapi3.Parameters{
			&openapi3.ParameterRef{
				Value: &openapi3.Parameter{
					Name:        "id",
					In:          openapi3.ParameterInPath,
					Description: fmt.Sprintf("ID of the %s", entity.Name),
					Required:    true,
					Schema: &openapi3.SchemaRef{
						Value: &openapi3.Schema{
							Type:   openapi3.TypeString,
							Format: "uuid",
						},
					},
				},
			},
		},
		RequestBody: &openapi3.RequestBodyRef{
			Value: &openapi3.RequestBody{
				Description: fmt.Sprintf("%s to update", entity.Name),
				Content: openapi3.Content{
					"application/json": {
						Schema: openapi3.NewSchemaRef(fmt.Sprintf("#/components/schemas/Update%s", entity.Name), nil),
					},
				},
				Required: true,
			},
		},
		Responses: openapi3.Responses{
			"200": &openapi3.ResponseRef{
				Value: &openapi3.Response{
					Description: lo.ToPtr(fmt.Sprintf("%supdated", entity.Name)),
					Content: openapi3.Content{
						"application/json": &openapi3.MediaType{
							Schema: &openapi3.SchemaRef{
								Ref: fmt.Sprintf("#/components/schemas/%s", entity.Name),
							},
						},
					},
				},
			},
			"400": &openapi3.ResponseRef{
				Ref: "#/components/responses/400",
			},
			"404": &openapi3.ResponseRef{
				Ref: "#/components/responses/404",
			},
			"409": &openapi3.ResponseRef{
				Ref: "#/components/responses/409",
			},
			"500": &openapi3.ResponseRef{
				Ref: "#/components/responses/500",
			},
			"default": &openapi3.ResponseRef{
				Ref: "#/components/responses/Error",
			},
		},
	}
}

func makeDelete(entity *entity.Entity) *openapi3.Operation {
	return &openapi3.Operation{
		Tags: []string{
			entity.Name,
		},
		Summary:     fmt.Sprintf("Delete %s", entity.Name),
		Description: fmt.Sprintf("Delete %s.", entity.Name),
		OperationID: fmt.Sprintf("delete%s", entity.Name),
		Parameters: openapi3.Parameters{
			&openapi3.ParameterRef{
				Value: &openapi3.Parameter{
					Name:        "id",
					In:          openapi3.ParameterInPath,
					Description: fmt.Sprintf("ID of the %s", entity.Name),
					Required:    true,
					Schema: &openapi3.SchemaRef{
						Value: &openapi3.Schema{
							Type:   openapi3.TypeString,
							Format: "uuid",
						},
					},
				},
			},
		},
		Responses: openapi3.Responses{
			"200": &openapi3.ResponseRef{
				Value: &openapi3.Response{
					Description: lo.ToPtr(fmt.Sprintf("%supdated", entity.Name)),
					Content: openapi3.Content{
						"application/json": &openapi3.MediaType{
							Schema: &openapi3.SchemaRef{
								Value: &openapi3.Schema{
									Type: openapi3.TypeObject,
									Properties: openapi3.Schemas{
										"status": &openapi3.SchemaRef{
											Value: &openapi3.Schema{
												Type: openapi3.TypeString,
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"400": &openapi3.ResponseRef{
				Ref: "#/components/responses/400",
			},
			"404": &openapi3.ResponseRef{
				Ref: "#/components/responses/404",
			},
			"409": &openapi3.ResponseRef{
				Ref: "#/components/responses/409",
			},
			"500": &openapi3.ResponseRef{
				Ref: "#/components/responses/500",
			},
			"default": &openapi3.ResponseRef{
				Ref: "#/components/responses/Error",
			},
		},
	}
}

func makeGetOne(entity *entity.Entity) *openapi3.Operation {
	params := openapi3.Parameters{
		&openapi3.ParameterRef{
			Value: &openapi3.Parameter{
				Name:        "id",
				In:          openapi3.ParameterInPath,
				Description: fmt.Sprintf("ID of the %s", entity.Name),
				Required:    true,
				Schema: &openapi3.SchemaRef{
					Value: &openapi3.Schema{
						Type:   openapi3.TypeString,
						Format: "uuid",
					},
				},
			},
		},
	}
	if withParam := makeWithParam(entity.Edges); withParam != nil {
		params = append(params, withParam)
	}
	return &openapi3.Operation{
		Tags: []string{
			entity.Name,
		},
		Summary:     fmt.Sprintf("Find a %s by ID", entity.Name),
		Description: fmt.Sprintf("Finds the %s with the requested ID and returns it.", entity.Name),
		OperationID: fmt.Sprintf("read%s", entity.Name),
		Parameters:  params,
		Responses: openapi3.Responses{
			"200": &openapi3.ResponseRef{
				Value: &openapi3.Response{
					Description: lo.ToPtr(fmt.Sprintf("%supdated", entity.Name)),
					Content: openapi3.Content{
						"application/json": &openapi3.MediaType{
							Schema: &openapi3.SchemaRef{
								Ref: fmt.Sprintf("#/components/schemas/%s", entity.Name),
							},
						},
					},
				},
			},
			"400": &openapi3.ResponseRef{
				Ref: "#/components/responses/400",
			},
			"404": &openapi3.ResponseRef{
				Ref: "#/components/responses/404",
			},
			"409": &openapi3.ResponseRef{
				Ref: "#/components/responses/409",
			},
			"500": &openapi3.ResponseRef{
				Ref: "#/components/responses/500",
			},
			"default": &openapi3.ResponseRef{
				Ref: "#/components/responses/Error",
			},
		},
	}
}

type getConfig struct {
	name string
}

func makeGet(cfg getConfig) *openapi3.Operation {
	return &openapi3.Operation{
		Tags: []string{
			cfg.name,
		},
		Summary:     fmt.Sprintf("Get %s", cfg.name),
		Description: fmt.Sprintf("Get %s.", cfg.name),
		OperationID: fmt.Sprintf("%s", cfg.name),
		Responses: openapi3.Responses{
			"200": &openapi3.ResponseRef{
				Value: &openapi3.Response{
					Description: lo.ToPtr(fmt.Sprintf("%supdated", cfg.name)),
					Headers: openapi3.Headers{
						"Content-Range": {
							Value: &openapi3.Header{
								Parameter: openapi3.Parameter{
									Required: true,
									Schema: &openapi3.SchemaRef{
										Value: &openapi3.Schema{
											Type:        openapi3.TypeInteger,
											Description: "Total items count",
										},
									},
								},
							},
						},
					},
					Content: openapi3.Content{
						"application/json": &openapi3.MediaType{
							Schema: &openapi3.SchemaRef{
								Ref: fmt.Sprintf("#/components/schemas/%s", cfg.name),
							},
						},
					},
				},
			},
			"400": &openapi3.ResponseRef{
				Ref: "#/components/responses/400",
			},
			"404": &openapi3.ResponseRef{
				Ref: "#/components/responses/404",
			},
			"409": &openapi3.ResponseRef{
				Ref: "#/components/responses/409",
			},
			"500": &openapi3.ResponseRef{
				Ref: "#/components/responses/500",
			},
			"default": &openapi3.ResponseRef{
				Ref: "#/components/responses/Error",
			},
		},
	}
}
