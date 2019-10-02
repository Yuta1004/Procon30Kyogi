package battle

import (
	"github.com/Yuta1004/procon30-kyogi/manager"
)

func getScore(battle manager.Battle) [][]int {
	// error check
	if len(battle.DetailInfo.Teams) <= 0 {
		return [][]int{[]int{0, 0}, []int{0, 0}}
	}

	scoreA := []int{
		battle.DetailInfo.Teams[0].AreaPoint, battle.DetailInfo.Teams[0].TilePoint,
	}
	scoreB := []int{
		battle.DetailInfo.Teams[1].AreaPoint, battle.DetailInfo.Teams[1].TilePoint,
	}
	myTeamID := battle.Info.TeamID
	if myTeamID == battle.DetailInfo.Teams[0].TeamID {
		return [][]int{scoreA, scoreB}
	}
	return [][]int{scoreA, scoreB}
}
