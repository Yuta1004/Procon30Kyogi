package solver

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Yuta1004/procon30-kyogi/config"
	"github.com/Yuta1004/procon30-kyogi/manager/battle"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"os"
	"strconv"
)

// ExecSolver : ソルバを起動, 実行する
func ExecSolver(ch chan string, battle battle.Battle) {
	// setting json
	jsonFName := strconv.Itoa(battle.Info.ID) + "_" + strconv.Itoa(battle.Turn)
	jsonStr, err := json.Marshal(battle.DetailInfo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not convert to json from \"battle\" : %s\n", err)
		ch <- "Error"
		return
	}
	saveJSON(jsonFName, jsonStr)

	// crate client
	client, err := client.NewClientWithOpts(client.WithVersion("1.40"))
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
	battleID := battle.Info.ID
	maxTurn := battle.Info.MaxTurn
	jsonIn := "/tmp/input.json"
	jsonOut := "/tmp/output.json"
	memLim := 999999999
	execTimeLim := int(float64(battle.Info.TurnMillis) * 0.7)
	confCont := container.Config{
		Image: image,
		Cmd: []string{
			"/bin/sh", "-c",
			fmt.Sprintf(
				"./solver.py %s %s %d %s %s %d %d %d && cat %s",
				jsonIn, jsonOut, battleID, "A", "B", maxTurn, execTimeLim, memLim, jsonOut,
			),
		},
		WorkingDir: "/tmp/",
		Tty:        true,
	}

	// config(host)
	confHost := container.HostConfig{
		AutoRemove: true,
		Mounts: []mount.Mount{
			{
				Type:   mount.TypeBind,
				Source: rootPath + "/tmp/" + jsonFName + ".json",
				Target: "/tmp/input.json",
			},
		},
	}

	// create
	ctx := context.Background()
	cont, err := client.ContainerCreate(ctx, &confCont, &confHost, &network.NetworkingConfig{}, "Procon30_"+jsonFName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not create container : %s\n", err)
		ch <- "Error"
		return
	}

	// exec -> attach -> remove
	client.ContainerStart(ctx, cont.ID, types.ContainerStartOptions{})
	ch <- "Success"
	return
}
