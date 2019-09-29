package config

import (
	"testing"
)

func TestGetConfigData(t *testing.T) {
	config := GetConfigData()
	if config == nil && config.GameServer.URL != "" && config.Solver.Image != "" {
		t.Errorf("Cannot read config.toml")
	}
}
