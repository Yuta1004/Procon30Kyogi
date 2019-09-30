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

var allBattleDict map[int]Battle

// BManagerExec : 名前の通り, 参加している試合全ての管理をする
func BManagerExec(token string) {

}

func makeAllBattleDict(token string) {
	battleInfoList := connector.GetAllBattle(token)
	for _, battleInfo := range *battleInfoList {
		battleDetailInfo := connector.GetBattleDetail(battleInfo.ID, token)
		allBattleDict[battleInfo.ID] = Battle{
			Info:       &battleInfo,
			DetailInfo: battleDetailInfo,
			Turn:       battleDetailInfo.Turn,
			SolverCh:   make(chan string, 10),
		}
	}
}
