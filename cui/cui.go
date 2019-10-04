package cui

import (
	"github.com/Yuta1004/procon30-kyogi/config"
	"github.com/Yuta1004/procon30-kyogi/manager/battle"
	"github.com/Yuta1004/procon30-kyogi/mylog"
	"github.com/eiannone/keyboard"
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
		if inp != rune('\n') {
			mylog.SetInputArea(string(inpBuf))
			continue
		}

		// command
		command := strings.Split(string(inpBuf), " ")
		switch command[0] {
		case "refresh":
			mylog.Warning("試合情報を再取得します...(更新終了まで操作をしないでください)")
			conf := config.GetConfigData()
			battle.MakeAllBattleDict(conf.GameServer.Token)

		case "exit":
			mylog.Info("システムを終了します...")
			os.Exit(0)
		}

		// clean buf
		inpBuf = make([]rune, 0)
		mylog.SetInputArea(string(inpBuf))
	}
}

func monitorStdin(ch chan rune) {
	keyboard.Open()
	defer keyboard.Close()

	for {
		char, key, _ := keyboard.GetKey()
		switch key {
		case keyboard.KeyBackspace, keyboard.KeyBackspace2:
			inpBuf = inpBuf[:max(len(inpBuf)-1, 0)]
		case keyboard.KeyEnter:
			char = '\n'
		case keyboard.KeySpace:
			char = ' '
			inpBuf = append(inpBuf, char)
		default:
			inpBuf = append(inpBuf, char)
		}
		ch <- char
	}
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
