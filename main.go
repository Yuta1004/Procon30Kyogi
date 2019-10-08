package main

import (
	"fmt"
	"github.com/Yuta1004/procon30-kyogi/config"
	"github.com/Yuta1004/procon30-kyogi/cui"
	"github.com/Yuta1004/procon30-kyogi/manager/battle"
	"github.com/Yuta1004/procon30-kyogi/mylog"
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
		mylog.Notify("BattleManagerを再起動します...")
	}
}
