package config

import (
	"testing"
)

func TestGetConfigData(t *testing.T) {
	config := GetConfigData()
	if config == nil && config.GameServer.URL != "" {
		t.Errorf("Cannot read config.toml")
	}
}
