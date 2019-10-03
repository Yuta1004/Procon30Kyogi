package cui

import (
	"github.com/Yuta1004/procon30-kyogi/mylog"
	"github.com/eiannone/keyboard"
	// "log"
	"os"
)

var inpBuf []rune

// CUI : cui
func CUI() {
	// init variables
	ch := make(chan rune)
	inpBuf = make([]rune, 1024)

	// start monitor
	go monitorStdin(ch)

	// mainloop
	for {
		// input
		inp := <-ch
		mylog.SetInputArea(string(inpBuf))
		if inp != rune('\n') {
			continue
		}

		// command
		mylog.Warning("-%s-", string(inpBuf))
		switch string(inpBuf) {
		case "exit":
			mylog.Info("システムを終了します...")
			os.Exit(0)
		}

		// clean buf
		inpBuf = nil
		inpBuf = make([]rune, 1024)
		mylog.SetInputArea(string(inpBuf))
	}
}

func monitorStdin(ch chan rune) {
	keyboard.Open()
	defer keyboard.Close()

	for {
		char, key, _ := keyboard.GetKey()
		switch key {
		case keyboard.KeyEsc:
			os.Exit(0)
		case keyboard.KeyBackspace, keyboard.KeyBackspace2:
			inpBuf = inpBuf[:len(inpBuf)-1]
		case keyboard.KeyEnter:
			inpBuf = append(inpBuf, 0)
			char = '\n'
		default:
			inpBuf = append(inpBuf, char)
		}
		ch <- char
	}
}
