package battle

import (
	"github.com/Yuta1004/procon30-kyogi/connector"
)

// Battle : 試合情報を扱う
type Battle struct {
	Info       *connector.BattleInfo
	DetailInfo *connector.BattleDetailInfo
	Turn       int
	SolverCh   chan string
}
