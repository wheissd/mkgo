---
sidebar_position: 0
---

# Getting started

MKGO is a tool aimed to reduce amount of work, required to start and maintain a go project.
It provides commands to initialize project, add models(with ent), auto-generate migrations(with atlas),
auto-generate crud operations with openapi schema.

1) install
1) initialize
1) describe your ent schema
1) generate
1) enjoy

## Prerequisites

mkgo has dependencies:
atlas, ent, ogen, goimports

```console
curl -sSf https://atlasgo.sh | ATLAS_VERSION=v0.19.3-cfa638c-canary sh
go install golang.org/x/tools/cmd/goimports@latest
``` 

## Installation

```console
go install github.com/wheissd/mkgo/cmd/mkgo
```

:::tip TIP

Ensure your PATH contains go bin dir

:::

## Initialize project

```console
mkgo init hello_mkgo
```

This creates basic project structure inside current folder

api
ent
openapi
gen_config.yaml
generate.go
```console
project_root
├── internal
│   └── api
│       ├── cmd
│       │   └── apigen
│       │       ├── main.go
│       │       └── pre_gen.go
│       └── gen
│           ├── cmd
│           │   └── main.go
│           ├── schema
│           └── generate.go
├── ent
├── openapi
├── gen_config.yaml
└── generate.go
```

## Create model

```console
mkgo model Example
```
Creates new ent model. For ent docs visit https://entgo.io/ 

## Generate crud handlers

After you described your entities, you can auto-generate crud operations for them.
There are two options:

1) Enable operations by default and disable them where you want.
2) Disable operations by default and enable them where you want.

:::tip TIP

By default cruds are not generated and fields will not be exposed to generated cruds

:::

### Generate crud operations by default and if you want to disable some for your entity - do it with annotations.
```yaml title='mkgo_config.yaml'
    - rest-client:
        OutputPath: api
        OpenApiPath: openapi/rest_client.json
        Mode: rest_client
        Title: Your API title
        Servers:
          - http://localhost:9000
        EnableDefaultReadOne: true
        EnableDefaultReadMany: true
        EnableDefaultCreate: true
        EnableDefaultUpdate: true
        EnableDefaultDelete: true
        FieldsPublicByDefault: true
```
```go title='your_model_schema.go' {13-28}
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/wheissd/gomk/annotations"
)

type Cat struct {
	ent.Schema
}

func (Cat) Annotations() []schema.Annotation {
	return []schema.Annotation{
		annotations.Entity().DisableDelete(annotations.Modes{"api"}),
	}
}

func (Cat) Fields() []ent.Field {
	return []ent.Field{
		// this field will be expose in api
		field.String("name"),
                // this field wont be exposed in api
		field.Int64("speed").Annotations(
			annotations.Field().SetPrivate(),
		),
	}
}

```

### Do not generate crud operations by default and if you want to enable some for your entity - do it with annotations.
:::tip TIP
You can leave config options empty - false is default value
:::
```yaml title='mkgo_config.yaml'
    - rest-client:
        OutputPath: api
        OpenApiPath: openapi/rest_client.json
        Mode: rest_client
        Title: Your API title
        Servers:
          - http://localhost:9000
```
```go title='your_model_schema.go' {13-28}
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/wheissd/gomk/annotations"
)

type Cat struct {
	ent.Schema
}

func (Cat) Annotations() []schema.Annotation {
	return []schema.Annotation{
		annotations.Entity().EnableCreate(annotations.Modes{"api"}),
	}
}

func (Cat) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Annotations(
			annotations.Field().SetPublic(),
		),
		field.Int64("speed").Annotations(
			annotations.Field().SetPublic(),
		),
	}
}

```
## Generate app
```console
mkgo generate
```
## Run
:::tip Tip

By default you'll have *rest_client* and *grpc_admin* as [your_generated_api_binary]

:::
```console
go run cmd/[your_generated_api_binary]
```
Voila, now you can test and deploy your api
