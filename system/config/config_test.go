package config

import "testing"

func TestConfig(t *testing.T) {
	cfg, viper, err := NewConfig("../../conf")
	if err != nil {
		t.Errorf("LoadConfig.Error: %#v", err)
	}

	if viper == nil {
		t.Errorf("Viper.Error: %#v", viper)
	}

	if cfg == nil {
		t.Errorf("Config.Error: %#v", cfg)
	}

	if cfg != nil {
		t.Logf("Config: %#v", cfg)
	}
}
