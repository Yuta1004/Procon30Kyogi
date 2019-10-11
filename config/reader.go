package config

import (
	"github.com/BurntSushi/toml"
	"github.com/Yuta1004/procon30-kyogi/mylog"
)

var (
	config Config
)

// Config : 設定情報を扱う構造体
type Config struct {
	GameServer GameServer
	Solver     Solver
}

// GameServer : 設定情報(GameServer)を扱う構造体
type GameServer struct {
	URL   string
	Token string
}

// Solver : 設定情報(Solver)を扱う構造体
type Solver struct {
	Image     string         // デフォルト(ファイルから読む)
	manualSet map[int]string // 試合IDごとにもつ
}

// GetConfigData : 設定情報を返す
func GetConfigData() *Config {
	if config.GameServer.URL == "" {
		_, err := toml.DecodeFile("config.toml", &config)
		config.Solver.manualSet = make(map[int]string)
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
	if config.Solver.manualSet == nil {
		config.Solver.manualSet = make(map[int]string)
	}
}

// Get : 任意の試合IDで使うソルバイメージを参照する
func (s Solver) Get(battleID int) string {
	val, exists := s.manualSet[battleID]
	if exists {
		return val
	}
	return s.Image
}

// Set : 任意の試合IDで使用するソルバイメージを変更する
func (s Solver) Set(battleID int, image string) {
	s.manualSet[battleID] = image
}
