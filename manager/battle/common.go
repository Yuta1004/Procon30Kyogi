package battle

import (
	"github.com/Yuta1004/procon30-kyogi/manager"
	"github.com/Yuta1004/procon30-kyogi/mylog"
	"time"
)

func getScore(battle manager.Battle) [][]int {
	// error check
	if len(battle.DetailInfo.Teams) <= 0 {
		return [][]int{[]int{0, 0}, []int{0, 0}}
	}

	scoreA := []int{battle.DetailInfo.Teams[0].AreaPoint, battle.DetailInfo.Teams[0].TilePoint}
	scoreB := []int{battle.DetailInfo.Teams[1].AreaPoint, battle.DetailInfo.Teams[1].TilePoint}
	myTeamID := battle.Info.TeamID
	if myTeamID == battle.DetailInfo.Teams[0].TeamID {
		return [][]int{scoreA, scoreB}
	}
	return [][]int{scoreB, scoreA}
}

func checkNeedUpdateBattle(battle manager.Battle) (bool, int) {
	turnMillis := battle.Info.IntervalMillis + battle.Info.TurnMillis
	nowUnix := int(time.Now().UnixNano() / 1000000)
	elapsedTime := nowUnix - battle.DetailInfo.StartedAtUnixTime*1000
	elapsedTurn := int(elapsedTime / turnMillis)
	return (0 <= elapsedTime && 0 <= elapsedTurn && elapsedTurn <= battle.Info.MaxTurn+1 && battle.Turn != elapsedTurn),
		elapsedTurn
}

func outBattleLog(battle manager.Battle) {
	score := getScore(battle)
	mylog.Notify("次ターンに移行しました -> BattleID: %d, Turn : %d", battle.Info.ID, battle.Turn)
	mylog.Info(
		"試合情報 -> \x1b[1mBattleID: %d, \x1b[34m自チーム: %d (A %d, T %d), \x1b[31m相手チーム: %d (A %d, T %d)",
		battle.Info.ID, score[0][0]+score[0][1], score[0][0], score[0][1], score[1][0]+score[1][1], score[1][0], score[1][1],
	)
}
