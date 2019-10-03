package cui

import (
	"github.com/Yuta1004/procon30-kyogi/mylog"
	"github.com/eiannone/keyboard"
	// "log"
	"os"
)

var inpBuf []byte

// CUI : cui
func CUI() {
	// init variables
	ch := make(chan byte)
	inpBuf = make([]byte, 1024)

	// start monitor
	go monitorStdin(ch)

	// mainloop
	for {
		// input
		inp := <-ch
		mylog.SetInputArea(string(inpBuf))
		if inp != byte('\n') {
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
		inpBuf = make([]byte, 1024)
	}
}

func monitorStdin(ch chan byte) {
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
			char = '\n'
		default:
			inpBuf = append(inpBuf, byte(char))
		}
		ch <- byte(char)
	}
}
