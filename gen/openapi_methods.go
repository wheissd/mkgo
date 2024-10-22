package gen

import (
	"fmt"
	"net/http"
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
						Type: &openapi3.Types{openapi3.TypeString},
					},
				},
			},
		}
	}
	return nil
}

func makeGetList(sch *entity.Schema, entityItem *entity.Entity) *openapi3.Operation {
	params := openapi3.Parameters{}
	// add filters params
	for _, f := range entityItem.Fields {
		t := f.Openapi.Type
		if t == "" {
			t = f.Type.Type.OApiTypeName()
		}
		if f.Type.Type.CanRange() {
			params = append(params, &openapi3.ParameterRef{
				Value: &openapi3.Parameter{
					In:          openapi3.ParameterInQuery,
					Name:        "filterRangeFrom" + cases.Pascal(f.Name),
					Description: "Range from filter by " + f.Name + " field",
					Schema: &openapi3.SchemaRef{
						Value: &openapi3.Schema{
							Type:   &openapi3.Types{t},
							Format: paramFormat(f),
						},
					},
				},
			}, &openapi3.ParameterRef{
				Value: &openapi3.Parameter{
					In:          openapi3.ParameterInQuery,
					Name:        "filterRangeTo" + cases.Pascal(f.Name),
					Description: "Range to filter by " + f.Name + " field",
					Schema: &openapi3.SchemaRef{
						Value: &openapi3.Schema{
							Type:   &openapi3.Types{t},
							Format: paramFormat(f),
						},
					},
				},
			})
		}
		params = append(params, &openapi3.ParameterRef{
			Value: &openapi3.Parameter{
				In:          openapi3.ParameterInQuery,
				Name:        "filter" + cases.Pascal(f.Name),
				Description: "Filter by " + f.Name + " field",
				Schema: &openapi3.SchemaRef{
					Value: &openapi3.Schema{
						Type:   &openapi3.Types{t},
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
						Type:   &openapi3.Types{openapi3.TypeString},
						Format: "uuid",
					},
				},
			},
		})
	}
	if withParam := makeWithParam(entityItem.Edges); withParam != nil {
		params = append(params, withParam)
	}

	responses := openapi3.NewResponses(
		openapi3.WithStatus(http.StatusOK, &openapi3.ResponseRef{
			Value: &openapi3.Response{
				Description: lo.ToPtr(fmt.Sprintf("result %s options headers", entityItem.Name)),
				Headers: openapi3.Headers{
					"Content-Range": {
						Value: &openapi3.Header{
							Parameter: openapi3.Parameter{
								Required: true,
								Schema: &openapi3.SchemaRef{
									Value: &openapi3.Schema{
										Type:        &openapi3.Types{openapi3.TypeInteger},
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
		}),
		openapi3.WithStatus(http.StatusBadRequest, &openapi3.ResponseRef{
			Ref: "#/components/responses/400",
		}),
		openapi3.WithStatus(http.StatusNotFound, &openapi3.ResponseRef{
			Ref: "#/components/responses/404",
		}),
		openapi3.WithStatus(http.StatusConflict, &openapi3.ResponseRef{
			Ref: "#/components/responses/409",
		}),
		openapi3.WithStatus(http.StatusInternalServerError, &openapi3.ResponseRef{
			Ref: "#/components/responses/500",
		}),
	)
	responses.Set("default", &openapi3.ResponseRef{
		Ref: "#/components/responses/Error",
	})
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
							Type: &openapi3.Types{openapi3.TypeInteger},
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
							Type: &openapi3.Types{openapi3.TypeInteger},
							Min:  lo.ToPtr(float64(1)),
							Max:  lo.ToPtr(float64(255)),
						},
					},
				},
			}}, params...),
		Responses: responses,
	}
}

func makeCreate(entity *entity.Entity) *openapi3.Operation {
	responses := openapi3.NewResponses()
	responses.Set("200", &openapi3.ResponseRef{
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
	})
	responses.Set("400", &openapi3.ResponseRef{
		Ref: "#/components/responses/400",
	})
	responses.Set("404", &openapi3.ResponseRef{
		Ref: "#/components/responses/404",
	})
	responses.Set("409", &openapi3.ResponseRef{
		Ref: "#/components/responses/409",
	})
	responses.Set("500", &openapi3.ResponseRef{
		Ref: "#/components/responses/500",
	})
	responses.Set("default", &openapi3.ResponseRef{
		Ref: "#/components/responses/Error",
	})
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
		Responses: responses,
	}
}

func makeUpdate(entity *entity.Entity) *openapi3.Operation {
	responses := openapi3.NewResponses()
	responses.Set("200", &openapi3.ResponseRef{
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
	})
	responses.Set("400", &openapi3.ResponseRef{
		Ref: "#/components/responses/400",
	})
	responses.Set("404", &openapi3.ResponseRef{
		Ref: "#/components/responses/404",
	})
	responses.Set("409", &openapi3.ResponseRef{
		Ref: "#/components/responses/409",
	})
	responses.Set("500", &openapi3.ResponseRef{
		Ref: "#/components/responses/500",
	})
	responses.Set("default", &openapi3.ResponseRef{
		Ref: "#/components/responses/Error",
	})
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
							Type:   &openapi3.Types{openapi3.TypeString},
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
		Responses: responses,
	}
}

func makeDelete(entity *entity.Entity) *openapi3.Operation {
	responses := openapi3.NewResponses()
	responses.Set("200", &openapi3.ResponseRef{
		Value: &openapi3.Response{
			Description: lo.ToPtr(fmt.Sprintf("%supdated", entity.Name)),
			Content: openapi3.Content{
				"application/json": &openapi3.MediaType{
					Schema: &openapi3.SchemaRef{
						Value: &openapi3.Schema{
							Type: &openapi3.Types{openapi3.TypeObject},
							Properties: openapi3.Schemas{
								"status": &openapi3.SchemaRef{
									Value: &openapi3.Schema{
										Type: &openapi3.Types{openapi3.TypeString},
									},
								},
							},
						},
					},
				},
			},
		},
	})
	responses.Set("400", &openapi3.ResponseRef{
		Ref: "#/components/responses/400",
	})
	responses.Set("404", &openapi3.ResponseRef{
		Ref: "#/components/responses/404",
	})
	responses.Set("409", &openapi3.ResponseRef{
		Ref: "#/components/responses/409",
	})
	responses.Set("500", &openapi3.ResponseRef{
		Ref: "#/components/responses/500",
	})
	responses.Set("default", &openapi3.ResponseRef{
		Ref: "#/components/responses/Error",
	})
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
							Type:   &openapi3.Types{openapi3.TypeString},
							Format: "uuid",
						},
					},
				},
			},
		},
		Responses: responses,
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
						Type:   &openapi3.Types{openapi3.TypeString},
						Format: "uuid",
					},
				},
			},
		},
	}
	if withParam := makeWithParam(entity.Edges); withParam != nil {
		params = append(params, withParam)
	}
	responses := openapi3.NewResponses()
	responses.Set("200", &openapi3.ResponseRef{
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
	})
	responses.Set("400", &openapi3.ResponseRef{
		Ref: "#/components/responses/400",
	})
	responses.Set("404", &openapi3.ResponseRef{
		Ref: "#/components/responses/404",
	})
	responses.Set("409", &openapi3.ResponseRef{
		Ref: "#/components/responses/409",
	})
	responses.Set("500", &openapi3.ResponseRef{
		Ref: "#/components/responses/500",
	})
	responses.Set("default", &openapi3.ResponseRef{
		Ref: "#/components/responses/Error",
	})
	return &openapi3.Operation{
		Tags: []string{
			entity.Name,
		},
		Summary:     fmt.Sprintf("Find a %s by ID", entity.Name),
		Description: fmt.Sprintf("Finds the %s with the requested ID and returns it.", entity.Name),
		OperationID: fmt.Sprintf("read%s", entity.Name),
		Parameters:  params,
		Responses:   responses,
	}
}

type getConfig struct {
	name string
}

func makeGet(cfg getConfig) *openapi3.Operation {
	responses := openapi3.NewResponses()
	responses.Set("200", &openapi3.ResponseRef{
		Value: &openapi3.Response{
			Description: lo.ToPtr(fmt.Sprintf("%supdated", cfg.name)),
			Headers: openapi3.Headers{
				"Content-Range": {
					Value: &openapi3.Header{
						Parameter: openapi3.Parameter{
							Required: true,
							Schema: &openapi3.SchemaRef{
								Value: &openapi3.Schema{
									Type:        &openapi3.Types{openapi3.TypeInteger},
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
	})
	responses.Set("400", &openapi3.ResponseRef{
		Ref: "#/components/responses/400",
	})
	responses.Set("404", &openapi3.ResponseRef{
		Ref: "#/components/responses/404",
	})
	responses.Set("409", &openapi3.ResponseRef{
		Ref: "#/components/responses/409",
	})
	responses.Set("500", &openapi3.ResponseRef{
		Ref: "#/components/responses/500",
	})
	responses.Set("default", &openapi3.ResponseRef{
		Ref: "#/components/responses/Error",
	})
	return &openapi3.Operation{
		Tags: []string{
			cfg.name,
		},
		Summary:     fmt.Sprintf("Get %s", cfg.name),
		Description: fmt.Sprintf("Get %s.", cfg.name),
		OperationID: fmt.Sprintf("%s", cfg.name),
		Responses:   responses,
	}
}
