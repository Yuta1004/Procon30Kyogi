package battle

import (
	"github.com/Yuta1004/procon30-kyogi/connector"
	"github.com/Yuta1004/procon30-kyogi/manager"
	"log"
)

func copyAllBattleDict() (tmp map[int]manager.Battle) {
	tmp = make(map[int]manager.Battle)
	for key, val := range allBattleDict {
		tmp[key] = val
	}
	return
}

func makeAllBattleDict(token string) {
	// error check
	log.Printf("[INFO] 参加している全て試合の情報を取得しています... -> Token: %s\n", token)
	battleInfoList := connector.GetAllBattle(token)
	if len(*battleInfoList) == 0 {
		log.Printf("\x1b[31m[ERROR] 参加している試合が存在しないか、情報の取得に失敗しました -> MAKEALLBATTLEDICT001\x1b[0m\n")
		return
	}

	// make allBattleDict
	for _, battleInfo := range *battleInfoList {
		battle := makeBattleStruct(token, battleInfo.ID)
		battle.Info = battleInfo
		allBattleDict[battleInfo.ID] = battle
	}
}

func makeBattleStruct(token string, battleID int) manager.Battle {
	// error check
	log.Printf("[INFO] 試合情報詳細を取得しています... -> Token: %s, BattleID: %d", token, battleID)
	battleDetailInfo := connector.GetBattleDetail(battleID, token)
	if battleDetailInfo.Width == 0 {
		log.Printf("\x1b[31m[ERROR] 試合情報詳細の取得に失敗しました -> MAKEBATTLESTRUCT001\x1b[0m\n")
		return manager.Battle{}
	}
	return manager.Battle{
		DetailInfo: battleDetailInfo,
		Turn:       battleDetailInfo.Turn,
		SolverCh:   make(chan string, 10),
	}
}
