package battle

import (
	"github.com/Yuta1004/procon30-kyogi/manager"
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

func calcTimeStatus(battle manager.Battle) (int, int) {
	turnMillis := battle.Info.IntervalMillis + battle.Info.TurnMillis
	nowUnix := int(time.Now().UnixNano() / 1000000)
	elapsedTime := nowUnix - battle.DetailInfo.StartedAtUnixTime*1000
	elapsedTurn := int(elapsedTime/turnMillis) + 1
	return elapsedTime, elapsedTurn
}
