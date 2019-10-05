package cui

import (
	"github.com/eiannone/keyboard"
)

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
