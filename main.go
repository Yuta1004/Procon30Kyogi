package main

import (
	"github.com/Yuta1004/procon30-kyogi/config"
	"github.com/Yuta1004/procon30-kyogi/manager/battle"
	"log"
)

func main() {
	conf := config.GetConfigData()
	for {
		battle.BManagerExec(conf.GameServer.Token)
		log.Printf("\x1b[32m[NOIFY] BattleManagerを再起動します...\x1b[0m\n")
	}
}
