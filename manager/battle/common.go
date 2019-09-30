package battle

import (
	"fmt"
	"github.com/Yuta1004/procon30-kyogi/connector"
	"github.com/Yuta1004/procon30-kyogi/manager"
	"os"
)

func copyAllBattleDict() (tmp map[int]manager.Battle) {
	tmp = make(map[int]manager.Battle)
	for key, val := range allBattleDict {
		tmp[key] = val
	}
	return
}

func makeAllBattleDict(token string) {
	battleInfoList := connector.GetAllBattle(token)
	for _, battleInfo := range *battleInfoList {
		battle := makeBattleStruct(token, battleInfo.ID)
		battle.Info = &battleInfo
		allBattleDict[battleInfo.ID] = battle
	}
}

func makeBattleStruct(token string, battleID int) manager.Battle {
	fmt.Fprintf(os.Stderr, "Get Data : %d\n", battleID)
	battleDetailInfo := connector.GetBattleDetail(battleID, token)
	return manager.Battle{
		DetailInfo: battleDetailInfo,
		Turn:       battleDetailInfo.Turn,
		SolverCh:   make(chan string, 10),
	}
}
