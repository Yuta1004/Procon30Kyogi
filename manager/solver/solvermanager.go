package solver

import (
	"encoding/json"
	"fmt"
	"github.com/Yuta1004/procon30-kyogi/config"
	"github.com/Yuta1004/procon30-kyogi/manager"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"os"
	"strconv"
)

// ExecSolver : ソルバを起動, 実行する
func ExecSolver(ch chan string, battle manager.Battle) {
	// setting json
	jsonFName := strconv.Itoa(battle.Info.ID) + "_" + strconv.Itoa(battle.Turn)
	jsonStr, err := json.Marshal(battle.DetailInfo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		ch <- "Error"
		return
	}
	saveJSON(jsonFName, jsonStr)

	// config
	conf := config.GetConfigData()
	image := conf.Solver.Image

	// config(container)
	/* TODO : 適切なメモリ量割り当て */
	myID, opponentID := getTeamIDs(battle)
	battleID := battle.Info.ID
	maxTurn := battle.Info.MaxTurn
	jsonIn := "/tmp/input.json"
	jsonOut := "/tmp/output.json"
	memLim := 999999999
	execTimeLim := int(float64(battle.Info.TurnMillis) * 0.0006)
	confCont := container.Config{
		Image: image,
		Cmd: []string{
			"/bin/sh", "-c",
			fmt.Sprintf(
				"echo \"{}\" > /tmp/output.json && timeout -s %d ./solver.py %s %s %d %d %d %d %d %d && cat %s",
				execTimeLim, jsonIn, jsonOut, battleID, myID, opponentID, maxTurn, execTimeLim*1000, memLim, jsonOut,
			),
		},
		WorkingDir: "/tmp/",
		Tty:        true,
	}

	// config(host)
	confHost := container.HostConfig{
		Mounts: []mount.Mount{
			{
				Type:   mount.TypeBind,
				Source: rootPath + "/tmp/" + jsonFName + ".json",
				Target: "/tmp/input.json",
			},
		},
	}

	ch <- callContainer(&confCont, &confHost, "Procon30_"+jsonFName)
	return
}
