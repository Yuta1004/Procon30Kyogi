package solver

import (
	"encoding/json"
	"fmt"
	"github.com/Yuta1004/procon30-kyogi/config"
	"github.com/Yuta1004/procon30-kyogi/manager/battle"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"os"
	"strconv"
)

// ExecSolver : ソルバを起動, 実行する
func ExecSolver(ch chan string, battle battle.Battle) {
	// setting json
	jsonStr, err := json.Marshal(battle.DetailInfo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not convert to json from \"battle\" : %s\n", err)
		ch <- "Error"
		return
	}
	_ = jsonStr

	// crate client
	client, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not create docker clinet : %s\n", err)
		ch <- "Error"
		return
	}
	_ = client

	// config
	conf := config.GetConfigData()
	image := conf.Solver.Image

	// config(container)
	battleIDStr := strconv.Itoa(battle.Info.ID)
	maxTurnStr := strconv.Itoa(battle.Info.MaxTurn)
	jsonInPath := "/usr/input.json"
	jsonOutPath := "/usr/output.json"
	execTimeLim := "50000" // ms
	memLim := "100000"     // kb
	confCont := container.Config{
		Image: image,
		Cmd:   []string{"./solver.py", jsonInPath, jsonOutPath, battleIDStr, "A", "B", maxTurnStr, execTimeLim, memLim},
	}
	_ = confCont
}
