package viewer

import (
	"github.com/Yuta1004/procon30-kyogi/manager"
)

func getTeamIDs(battle manager.Battle) (int, int) {
	// size check
	if len(battle.DetailInfo.Teams) == 0 {
		return -1, -1
	}

	myTeamID := battle.Info.TeamID
	if battle.DetailInfo.Teams[0].TeamID == myTeamID {
		return myTeamID, battle.DetailInfo.Teams[1].TeamID
	}
	return myTeamID, battle.DetailInfo.Teams[0].TeamID
}
