package battle

import (
	"github.com/Yuta1004/procon30-kyogi/connector"
	"github.com/Yuta1004/procon30-kyogi/manager"
	"github.com/Yuta1004/procon30-kyogi/mylog"
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
	// error check
	mylog.Info("参加している全て試合の情報を取得しています... -> Token: %s", token)
	battleInfoList := connector.GetAllBattle(token)
	if len(*battleInfoList) == 0 {
		mylog.Error("参加している試合が存在しないか、情報の取得に失敗しました -> MAKEALLBATTLEDICT001")
		mylog.Error("システムを終了します...")
		os.Exit(1)
	}

	// make allBattleDict
	for _, battleInfo := range *battleInfoList {
		battle := makeBattleStruct(token, battleInfo.ID)
		battle.Info = battleInfo
		allBattleDict[battleInfo.ID] = battle
		mylog.Info("試合管理を始めます -> BattleID: %d", battle.Info.ID)
	}
}

func makeBattleStruct(token string, battleID int) manager.Battle {
	// error check
	mylog.Info("試合情報詳細を取得しています... -> Token: %s, BattleID: %d", token, battleID)
	battleDetailInfo := connector.GetBattleDetail(battleID, token)
	if battleDetailInfo.Width == 0 {
		mylog.Error("試合情報詳細の取得に失敗しました -> MAKEBATTLESTRUCT001")
		return manager.Battle{}
	}
	return manager.Battle{
		DetailInfo: battleDetailInfo,
		Turn:       battleDetailInfo.Turn,
		SolverCh:   make(chan string, 10),
	}
}
