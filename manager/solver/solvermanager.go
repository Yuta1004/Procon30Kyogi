package solver

import (
	"encoding/json"
	"fmt"
	"github.com/Yuta1004/procon30-kyogi/config"
	"github.com/Yuta1004/procon30-kyogi/manager/battle"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
	"os"
	"strconv"
)

// ExecSolver : ソルバを起動, 実行する
func ExecSolver(ch chan string, battle battle.Battle) {
	// setting json
	jsonFName := strconv.Itoa(battle.Info.ID) + strconv.Itoa(battle.Turn)
	jsonStr, err := json.Marshal(battle.DetailInfo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not convert to json from \"battle\" : %s\n", err)
		ch <- "Error"
		return
	}
	saveJSON(jsonFName, jsonStr)

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
	battleID := strconv.Itoa(battle.Info.ID)
	maxTurn := strconv.Itoa(battle.Info.MaxTurn)
	execTimeLim := strconv.Itoa(int(float64(battle.Info.TurnMillis) * 0.7))
	jsonInPath := "/usr/input.json"
	jsonOutPath := "/usr/output.json"
	memLim := "999999999"
	confCont := container.Config{
		Image: image,
		Cmd:   []string{"./solver.py", jsonInPath, jsonOutPath, battleID, "A", "B", maxTurn, execTimeLim, memLim},
	}
	_ = confCont

	// config(host)
	confHost := container.HostConfig{
		AutoRemove: true,
		Mounts: []mount.Mount{
			{
				Type:   mount.TypeBind,
				Source: rootPath + "/" + jsonFName + ".json",
				Target: "/usr/input.json",
			},
		},
	}
	_ = confHost
}
