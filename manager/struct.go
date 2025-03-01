package manager

import (
	"github.com/Yuta1004/procon30-kyogi/connector"
)

// Battle : 試合情報を扱う
type Battle struct {
	ProcessErrCnt int
	Info          connector.BattleInfo
	DetailInfo    connector.BattleDetailInfo
	Turn          int
	SolverCh      chan string
}
