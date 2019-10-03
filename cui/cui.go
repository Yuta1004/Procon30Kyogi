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
		<-ch
		mylog.SetInputArea(string(inpBuf))
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
		default:
			inpBuf = append(inpBuf, byte(char))
		}
		ch <- byte(char)
	}
}
