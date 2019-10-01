package solver

import (
	"github.com/Yuta1004/procon30-kyogi/manager"
	"log"
	"os"
	"path"
)

var (
	rootPath = path.Join(os.Getenv("GOPATH"), "src/github.com/Yuta1004/procon30-kyogi")
)

func saveJSON(name string, jsonBody []byte) bool {
	// mkdir
	_ = os.Mkdir(rootPath+"/tmp", 0744)

	// save json
	file, err := os.Create(rootPath + "/tmp/" + name + ".json")
	if err != nil {
		log.Printf("\x1b[31mソルバ起動準備中にエラーが発生しました -> SAVEJSON\x1b[0m\n")
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
