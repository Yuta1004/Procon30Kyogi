package config

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
	ManualSet map[int]string // 試合IDごとにもつ
}
