package noent

//go:generate go run internal/examples/noent/api/cmd/apigen -cfg_path=api_gen_config.yaml -ent_path internal/examples/noent/ent/gen
//go:generate go run github.com/ogen-go/ogen/cmd/ogen@latest --target api/gen --clean openapi/api.json --convenient-errors on
