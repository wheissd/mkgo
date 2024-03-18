package new2

//go:generate go run github.com/wheissd/gomk/internal/examples/new2/api/cmd/apigen -cfg_path=gen_config.yaml -ent_path internal/examples/new2/ent/gen
//go:generate go run github.com/ogen-go/ogen/cmd/ogen@latest --package ogen --target api/ogen --clean openapi/api.json --convenient-errors on
