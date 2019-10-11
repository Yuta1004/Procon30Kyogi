package config

import (
	"github.com/BurntSushi/toml"
	"github.com/Yuta1004/procon30-kyogi/mylog"
)

var (
	config Config
)

// GetConfigData : 設定情報を返す
func GetConfigData() *Config {
	if config.GameServer.URL == "" {
		_, err := toml.DecodeFile("config.toml", &config)
		config.Solver.ManualSet = make(map[int]string)
		if err != nil {
			mylog.Error("設定ファイルの読み込み中にエラーが発生しました")
			mylog.Error(err.Error())
			return nil
		}
	}
	return &config
}

// SetConfigData : 設定情報を新しくセットする
func SetConfigData(conf Config) {
	config = conf
	if config.Solver.ManualSet == nil {
		config.Solver.ManualSet = make(map[int]string)
	}
}
