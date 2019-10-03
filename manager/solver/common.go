package solver

import (
	"github.com/Yuta1004/procon30-kyogi/manager"
	"github.com/Yuta1004/procon30-kyogi/mylog"
	"os"
)

func saveJSON(name string, jsonBody []byte) bool {
	// mkdir
	_ = os.Mkdir("tmp", 0744)

	// save json
	file, err := os.Create("./tmp/" + name + ".json")
	if err != nil {
		mylog.Error("ソルバ起動準備中にエラーが発生しました -> SAVEJSON")
		return false
	}
	file.Write(jsonBody)
	file.Close()
	return true
}

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
