package config

import (
	"os"
	"strings"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

type parseConfig struct {
	path string
}

type Option func(config *parseConfig)

func OptionPath(path string) Option {
	return func(cfg *parseConfig) {
		cfg.path = path
	}
}

func Parse[CFG interface{}](cfg CFG, options ...Option) {
	parseCfg := parseConfig{
		path: "",
	}
	for _, opt := range options {
		opt(&parseCfg)
	}

	if parseCfg.path == "" {
		parseCfg.path = "config.yaml"
		args := os.Args
		for _, arg := range args {
			if strings.Contains(arg, "config") {
				splitted := strings.Split(arg, "=")
				if len(splitted) > 1 {
					parseCfg.path = splitted[1]
				}
			}
		}
	}

	k := koanf.New(".")
	if err := k.Load(file.Provider(parseCfg.path), yaml.Parser()); err != nil {
		panic(err)
	}

	err := k.Unmarshal("", cfg)
	if err != nil {
		panic(err)
	}
}
