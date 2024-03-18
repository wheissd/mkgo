package config

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

type Config struct {
	Env         string
	SecondLevel SecondLevel
}

type SecondLevel struct {
	ThirdLevel string
}

func TestConfig(t *testing.T) {
	var cfg Config
	Parse(&cfg)
	assert.Equal(t, cfg.Env, "local")
	assert.Equal(t, cfg.SecondLevel.ThirdLevel, "third")
}
