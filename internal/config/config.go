package config

import (
	genconfig "github.com/wheissd/mkgo/gen/config"
)

const DefaultMode = "default"

type GenConfig struct {
	APIs map[string]GenConfigItem `json:"apis"`
}

type GenConfigItem struct {
	OutputPath  string
	OpenApiPath string
	ProtoPath   string
	Mode        string
	Transport   string
	Title       string
	Version     string
	Servers     []string
	// enable create by default,
	// you can enable/disable explicitly for each model in genOpenapi annotations
	EnableDefaultCreate bool
	// enable update by default,
	// you can enable/disable explicitly for each model in genOpenapi annotations
	EnableDefaultUpdate bool
	// enable delete by default,
	// you can enable/disable explicitly for each model in genOpenapi annotations
	EnableDefaultDelete bool
	// enable readOne by default
	// you can enable/disable explicitly for each model in genOpenapi annotations
	EnableDefaultReadOne bool
	// enable readMany by default
	// you can enable/disable explicitly for each model in genOpenapi annotations
	EnableDefaultReadMany bool
	// make fields public by default, set field annotations to hide
	FieldsPublicByDefault bool
	// enable edges read by default, set edge annotations to hide
	EnableEdgeReadByDefault bool
	// enable edges write by default, set edge annotations to hide
	EnableEdgeCreateByDefault bool
	EnableEdgeUpdateByDefault bool
	EnableEdgeDeleteByDefault bool
	// enable filtering by default, set field annotations to hide
	EnableFilterByDefault bool
	// max level of with nesting in get(one/list) handlers, 0 = not limited
	WithMaxNesting int
}

func (c *GenConfig) Parse(path string) {
	genconfig.Parse(&c, genconfig.OptionPath(path))
	for i := range c.APIs {
		item := c.APIs[i]
		if c.APIs[i].Mode == "" {
			item.Mode = DefaultMode
		}
		if c.APIs[i].Transport == "" {
			item.Transport = "http"
		}
		if c.APIs[i].OpenApiPath == "" {
			item.OpenApiPath = "openapi/api.json"
		}
		if c.APIs[i].ProtoPath == "" {
			item.ProtoPath = "proto"
		}
		c.APIs[i] = item
	}
}
