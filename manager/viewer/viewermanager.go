package viewer

import (
	"github.com/Yuta1004/procon30-kyogi/config"
	"github.com/Yuta1004/procon30-kyogi/manager/battle"
	"github.com/Yuta1004/procon30-kyogi/mylog"
	"os/exec"
	"strconv"
)

// ExecViewer : ビューワ実行
func ExecViewer(battleID int) {
	// get data
	conf := config.GetConfigData()
	serverURL := conf.GameServer.URL
	token := conf.GameServer.Token
	allBattleDict := battle.GetBattleData()
	battle, ok := allBattleDict[battleID]
	if !ok {
		mylog.Error("無効な試合IDが指定されました -> BattleID: %d", battleID)
		return
	}

	// exec viewer
	myID, oppoID := getTeamIDs(battle)
	cmd := exec.Command(
		"python3",
		"./viewer/matchViewer.py",
		strconv.Itoa(battle.Info.ID),
		token,
		strconv.Itoa(myID),
		strconv.Itoa(oppoID),
		strconv.Itoa(battle.DetailInfo.StartedAtUnixTime),
		strconv.Itoa(battle.Info.MaxTurn),
		strconv.Itoa(battle.Info.TurnMillis),
		strconv.Itoa(battle.Info.IntervalMillis),
		serverURL,
	)
	if err := cmd.Start(); err != nil {
		mylog.Error("ビューワの起動に失敗しました")
		mylog.Error(err.Error())
		return
	}
	cmd.Wait()
	mylog.Notify("ビューワの実行を終了しました")
	defer cmd.Process.Kill()
}
