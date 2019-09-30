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

var battleList []Battle

// BManagerExec : 名前の通り, 参加している試合全ての管理をする
func BManagerExec() {

}
