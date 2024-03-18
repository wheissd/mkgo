package new3

//go:generate go run github.com/wheissd/gomk/internal/examples/new3/api/cmd/apigen -cfg_path=gen_config.yaml -ent_path internal/examples/new3/ent/gen
//go:generate go run github.com/ogen-go/ogen/cmd/ogen@latest --target api/gen --clean openapi/api.json --convenient-errors on
//go:generate protoc --proto_path proto --go_out=. --go-grpc_out=. --go-grpc_opt=paths=import human_gen.proto
