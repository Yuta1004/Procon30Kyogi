package main

import (
	"github.com/Yuta1004/procon30-kyogi/config"
	"github.com/Yuta1004/procon30-kyogi/manager/battle"
)

func main() {
	conf := config.GetConfigData()
	battle.BManagerExec(conf.GameServer.Token)
}
