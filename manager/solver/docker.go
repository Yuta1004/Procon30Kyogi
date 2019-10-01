package solver

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"io/ioutil"
	"log"
)

func callContainer(confCont *container.Config, confHost *container.HostConfig, name string) string {
	// crate client
	client, err := client.NewClientWithOpts(client.WithVersion("1.40"))
	if err != nil {
		log.Printf("\x1b[31m[ERROR] ソルバ起動中にエラーが発生しました -> CALLCONTAINER001\x1b[0m\n")
		return "{}"
	}

	// create
	ctx := context.Background()
	cont, err := client.ContainerCreate(ctx, confCont, confHost, &network.NetworkingConfig{}, name)
	if err != nil {
		log.Printf("\x1b[31m[ERROR] ソルバ起動中にエラーが発生しました -> CALLCONTAINER002\x1b[0m\n")
		return "{}"
	}

	// start
	err = client.ContainerStart(ctx, cont.ID, types.ContainerStartOptions{})
	if err != nil {
		log.Printf("\x1b[31m[ERROR] ソルバ起動中にエラーが発生しました -> CALLCONTAINER003\x1b[0m\n")
		return "{}"
	}
	defer client.ContainerRemove(ctx, cont.ID, types.ContainerRemoveOptions{})

	// wait...
	statusCh, errCh := client.ContainerWait(ctx, cont.ID, container.WaitConditionNotRunning)
	select {
	case <-errCh:
		log.Printf("\x1b[31m[ERROR] ソルバ起動中にエラーが発生しました -> CALLCONTAINER004\x1b[0m\n")
		return "{}"
	case <-statusCh:
	}

	// get log
	out, err := client.ContainerLogs(ctx, cont.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		log.Printf("\x1b[31m[ERROR] ソルバ実行中にエラーが発生しました -> CALLCONTAINER005\x1b[0m\n")
		return "{}"
	}
	result, _ := ioutil.ReadAll(out)
	return string(result)
}
