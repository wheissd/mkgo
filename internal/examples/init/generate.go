package init

//go:generate go run internal/examples/init/api/cmd/apigen -cfg_path=api_gen_config.yaml -ent_path internal/examples/init/ent/gen
//go:generate go run github.com/ogen-go/ogen/cmd/ogen@latest --target api/gen --clean openapi/api.json --convenient-errors on
