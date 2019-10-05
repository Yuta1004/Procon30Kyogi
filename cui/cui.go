package cui

import (
	"github.com/Yuta1004/procon30-kyogi/config"
	"github.com/Yuta1004/procon30-kyogi/manager/battle"
	"github.com/Yuta1004/procon30-kyogi/mylog"
	"os"
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
	case "viewer":
		if len(command) < 2 {
			mylog.Warning("Usage : viewer <BattleID>")
			return
		}
		mylog.Info("ビューワを起動します... -> BattleID : %s", command[1])

	case "solver":
		if len(command) < 2 {
			mylog.Warning("Usage : solver <SolverImage>")
			return
		}
		conf := config.GetConfigData()
		conf.Solver.Image = command[1]
		config.SetConfigData(*conf)
		mylog.Notify("使用するソルバイメージを変更しました -> %s", command[1])

	case "token":
		if len(command) < 2 {
			mylog.Warning("Usage : token <Token>")
			return
		}
		conf := config.GetConfigData()
		conf.GameServer.Token = command[1]
		config.SetConfigData(*conf)
		mylog.Notify("使用するトークンを変更しました -> %s", command[1])

	case "refresh":
		mylog.Warning("試合情報を再取得します...(更新終了まで操作をしないでください)")
		conf := config.GetConfigData()
		battle.MakeAllBattleDict(conf.GameServer.Token)

	case "exit":
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
