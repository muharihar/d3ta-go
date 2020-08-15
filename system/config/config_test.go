package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	cfg, viper, err := NewConfig("../../conf")
	if assert.NoError(t, err, "Error while creating config: NewConfig") {
		assert.NotNil(t, viper)
		assert.NotNil(t, cfg)
	}
}
