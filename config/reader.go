package config

import (
	"github.com/BurntSushi/toml"
	"os"
	"path"
)

var (
	config   Config
	rootPath = path.Join(os.Getenv("GOPATH"), "src/github.com/Yuta1004/procon30-kyogi")
)

// Config : 設定情報を扱う構造体
type Config struct {
	GameServer GameServer
}

// GameServer : 設定情報(GameServer)を扱う構造体
type GameServer struct {
	URL string
}

// GetConfigData : 設定情報を返す
func GetConfigData() *Config {
	if config.GameServer.URL == "" {
		_, err := toml.DecodeFile(rootPath+"/config.toml", &config)
		if err != nil {
			return nil
		}
	}
	return &config
}
