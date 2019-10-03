package main

import (
	"fmt"
	"github.com/Yuta1004/procon30-kyogi/config"
	"github.com/Yuta1004/procon30-kyogi/cui"
	"github.com/Yuta1004/procon30-kyogi/manager/battle"
	"log"
)

func main() {
	// setting fot mylog
	for cnt := 0; cnt < 3; cnt++ {
		fmt.Println()
	}

	// exec CUI
	go cui.CUI()

	// exec battle manager
	conf := config.GetConfigData()
	for {
		battle.BManagerExec(conf.GameServer.Token)
		log.Printf("\x1b[32m[NOIFY] BattleManagerを再起動します...\x1b[0m\n")
	}
}
