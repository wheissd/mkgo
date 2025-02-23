package gen

import (
	"fmt"
	"log/slog"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/samber/lo"
	"github.com/wheissd/mkgo/internal/config"
	"github.com/wheissd/mkgo/internal/entity"
	"github.com/wheissd/mkgo/options"
)

func genOpenapi(logger *slog.Logger, sch *entity.Schema, cfg config.GenConfigItem, externalConfig ...options.OpenapiSchemaOption) openapi3.T {
	paths := openapi3.Paths{}
	schemas := openapi3.Schemas{}
	for _, entity := range sch.Entities {
		logger.Debug("make " + entity.Name + " openapi schema")
		var (
			create   *openapi3.Operation
			put      *openapi3.Operation
			del      *openapi3.Operation
			getList  *openapi3.Operation
			getOne   *openapi3.Operation
			optsOne  *openapi3.Operation
			optsMany *openapi3.Operation
		)
		if needCreateOp(sch, entity) {
			create = makeCreate(entity)
			//optsMany = makeListOptions(entity)
		}
		if needUpdateOp(sch, entity) {
			put = makeUpdate(entity)
			//optsOne = makeEntOptions(entity)
		}
		if needDeleteOp(sch, entity) {
			del = makeDelete(entity)
			//optsOne = makeEntOptions(entity)
		}
		if needReadOneOp(sch, entity) {
			logger.Debug("generate readOne op")
			getOne = makeGetOne(entity)
			//optsOne = makeEntOptions(entity)
		}
		if needReadManyOp(sch, entity) {
			getList = makeGetList(sch, entity)
			//optsMany = makeListOptions(entity)
		}
		paths.Set("/"+entity.Path, &openapi3.PathItem{
			Get:     getList,
			Options: optsMany,
			Post:    create,
		})
		paths.Set("/"+entity.Path+"/{id}", &openapi3.PathItem{
			Put:     put,
			Options: optsOne,
			Get:     getOne,
			Delete:  del,
		})
		if needReadManyOp(sch, entity) {
			schemas[fmt.Sprintf("%sList", entity.Name)] = &openapi3.SchemaRef{
				Value: &openapi3.Schema{
					Type:  &openapi3.Types{openapi3.TypeArray},
					Items: openapi3.NewSchemaRef(fmt.Sprintf("#/components/schemas/%s", entity.Name), &openapi3.Schema{}),
				},
			}
		}

		if needEntity(sch, entity) {
			schemas[fmt.Sprintf("%s", entity.Name)] = &openapi3.SchemaRef{
				Value: &openapi3.Schema{
					Type:       &openapi3.Types{openapi3.TypeObject},
					Properties: readProperties(logger, sch, entity),
					Required:   readPropertiesRequired(entity),
				},
			}
		}
		if needCreateEntity(sch, entity) {
			schemas[fmt.Sprintf("Create%s", entity.Name)] = &openapi3.SchemaRef{
				Value: &openapi3.Schema{
					Type:       &openapi3.Types{openapi3.TypeObject},
					Properties: createProperties(entity),
					Required:   createPropertiesRequired(entity),
				},
			}
		}
		if needUpdateEntity(sch, entity) {
			schemas[fmt.Sprintf("Update%s", entity.Name)] = &openapi3.SchemaRef{
				Value: &openapi3.Schema{
					Type:       &openapi3.Types{openapi3.TypeObject},
					Properties: updateProperties(entity),
				},
			}
		}
	}
	schemas["Error"] = &openapi3.SchemaRef{
		Value: &openapi3.Schema{
			Type: &openapi3.Types{openapi3.TypeObject},
			Properties: openapi3.Schemas{
				"code": &openapi3.SchemaRef{
					Value: &openapi3.Schema{
						Type:   &openapi3.Types{openapi3.TypeInteger},
						Format: "int64",
					},
				},
				"message": &openapi3.SchemaRef{
					Value: &openapi3.Schema{
						Type: &openapi3.Types{openapi3.TypeString},
					},
				},
			},
		},
	}
	servers := openapi3.Servers{}
	for _, url := range cfg.Servers {
		servers = append(servers, &openapi3.Server{URL: url})
	}
	root := openapi3.T{
		OpenAPI: "3.0.3",
		Info: &openapi3.Info{
			Title:       cfg.Title,
			Description: cfg.Title + " api",
			Version:     "0.0.1",
		},
		Servers: servers,
		Paths:   &paths,
		Components: &openapi3.Components{
			Schemas: schemas,
			Responses: openapi3.ResponseBodies{
				"400":   &errorResponse,
				"403":   &errorResponse,
				"404":   &errorResponse,
				"409":   &errorResponse,
				"500":   &errorResponse,
				"Error": &errorResponse,
			},
		},
	}
	schemaExternal(&root, externalConfig...)
	return root
}

var (
	errorResponse = openapi3.ResponseRef{
		Value: &openapi3.Response{
			Description: lo.ToPtr("Error response"),
			Content: openapi3.Content{
				"application/json": &openapi3.MediaType{
					Schema: &openapi3.SchemaRef{
						Value: &openapi3.Schema{
							Type: &openapi3.Types{openapi3.TypeArray},
							Items: &openapi3.SchemaRef{
								Value: &openapi3.Schema{
									Type: &openapi3.Types{openapi3.TypeObject},
									Properties: openapi3.Schemas{
										"code": &openapi3.SchemaRef{
											Value: &openapi3.Schema{
												Type: &openapi3.Types{openapi3.TypeInteger},
											},
										},
										"message": &openapi3.SchemaRef{
											Value: &openapi3.Schema{
												Type: &openapi3.Types{openapi3.TypeString},
											},
										},
									},
									Required: []string{
										"code",
										"status",
										"message",
									},
								},
							},
						},
					},
				},
			},
		},
	}
)
