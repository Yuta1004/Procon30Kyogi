package cui

import (
	"github.com/Yuta1004/procon30-kyogi/config"
	"github.com/Yuta1004/procon30-kyogi/connector"
	"github.com/Yuta1004/procon30-kyogi/manager/battle"
	"github.com/Yuta1004/procon30-kyogi/manager/viewer"
	"github.com/Yuta1004/procon30-kyogi/mylog"
	"os"
	"strconv"
	"strings"
)

var inpBuf []rune

// CUI : cui
func CUI() {
	// init variables
	ch := make(chan rune)
	inpBuf = make([]rune, 0)
	go monitorStdin(ch)

	// mainloop
	for {
		// input
		inp := <-ch
		if inp != rune('\n') || len(inpBuf) == 0 {
			mylog.SetInputArea(string(inpBuf))
			continue
		}

		// command
		command := strings.Split(string(inpBuf), " ")
		execCommand(command...)

		// clean buf
		inpBuf = make([]rune, 0)
		mylog.SetInputArea(string(inpBuf))
	}
}

func execCommand(command ...string) {
	switch command[0] {
	case "viewer", "v":
		if len(command) < 2 {
			mylog.Warning("Usage : viewer <BattleID>")
			return
		}
		battleID, err := strconv.Atoi(command[1])
		if err != nil {
			mylog.Error("無効な試合IDが指定されました -> BattleID : %s", command[1])
			return
		}
		mylog.Notify("ビューワを起動します... -> BattleID : %s", command[1])
		go viewer.ExecViewer(battleID)

	case "solver", "s":
		if len(command) < 3 {
			mylog.Warning("Usage : solver <BattleID> <SolverVersion>")
			return
		}
		battleID, err := strconv.Atoi(command[1])
		if err != nil {
			mylog.Error("無効な試合IDが指定されました")
			return
		}
		conf := config.GetConfigData()
		conf.Solver.Set(battleID, "procon30-solver:Ver"+command[2])
		config.SetConfigData(*conf)
		mylog.Notify("使用するソルバイメージを変更しました -> BattleID: %s, SolverImage: procon30-solver:Ver%s", command[1], command[2])

	case "token":
		if len(command) < 2 {
			mylog.Warning("Usage : token <Token>")
			return
		}
		conf := config.GetConfigData()
		conf.GameServer.Token = command[1]
		config.SetConfigData(*conf)
		mylog.Notify("使用するトークンを変更しました -> %s", command[1])

	case "check":
		conf := config.GetConfigData()
		connector.CheckToken(conf.GameServer.Token)

	case "status":
		mylog.Notify("\x1b[1m----- 現在の試合状況 ----")
		for id, battle := range battle.GetBattleData() {
			mylog.Notify(
				"BattleID: %d, Turn: %d, MaxTurn: %d, MatchTo: %s",
				id, battle.Turn, battle.Info.MaxTurn, battle.Info.MatchTo,
			)
		}
		mylog.Notify("\x1b[1m-----------------------")

	case "config":
		config := config.GetConfigData()
		mylog.Notify("\x1b[1m----- 現在の設定状況 -----")
		mylog.Notify("ゲームサーバURL: %s", config.GameServer.URL)
		mylog.Notify("トークン: %s", config.GameServer.Token)
		mylog.Notify("ソルバイメージ: %s", config.Solver.Image)
		mylog.Notify("\x1b[1m------------------------")

	case "refresh":
		mylog.Warning("試合情報を再取得します...(更新終了まで操作をしないでください)")
		conf := config.GetConfigData()
		battle.MakeAllBattleDict(conf.GameServer.Token)

	case "exit", "q":
		mylog.Info("システムを終了します...")
		os.Exit(0)

	default:
		mylog.Warning("定義されていないコマンドです -> %s", command[0])
	}
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
