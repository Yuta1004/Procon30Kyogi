package config

// Config : 設定情報を扱う構造体
type Config struct {
	GameServer GameServer
}

// GameServer : 設定情報(GameServer)を扱う構造体
type GameServer struct {
	URL string
}
