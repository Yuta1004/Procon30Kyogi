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
	if config == nil && config.GameServer.URL != "" && config.Solver.Get(0) != "" {
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

func TestControlSolverConfig(t *testing.T) {
	// read 1
	os.Chdir(rootPath)
	SetConfigData(Config{GameServer: GameServer{URL: "test"}, Solver: Solver{Image: "test"}})
	conf := GetConfigData()
	if conf.Solver.Get(0) != "test" {
		t.Fatal("TestContronSolverConfig001")
	}

	// write1
	conf.Solver.Set(1204, "1204image")
	SetConfigData(*conf)
	conf = GetConfigData()
	if conf.Solver.Get(1204) != "1204image" || conf.Solver.Get(0) != "test" {
		t.Fatal("TestControlSolverConfig002")
	}
}
