package config

import (
	"os"
	"path"
	"testing"
)

var (
	rootPath = path.Join(os.Getenv("GOPATH"), "src/github.com/Yuta1004/procon30-kyogi")
)

func TestGetConfigData(t *testing.T) {
	os.Chdir(rootPath)
	config := GetConfigData()
	if config == nil && config.GameServer.URL != "" && config.Solver.Image != "" {
		t.Errorf("Cannot read config.toml")
	}
}

func TestSetConfigData(t *testing.T) {
	os.Chdir(rootPath)
	SetConfigData(Config{GameServer: GameServer{URL: "test"}})
	config := GetConfigData()
	if config.GameServer.URL != "test" {
		t.Fail()
	}
}
