package solver

import (
	"context"
	"github.com/Yuta1004/procon30-kyogi/mylog"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"io/ioutil"
)

func callContainer(confCont *container.Config, confHost *container.HostConfig, name string) string {
	// crate client
	client, err := client.NewClientWithOpts(client.WithVersion("1.40"))
	if err != nil {
		mylog.Error("ソルバ起動中にエラーが発生しました -> CALLCONTAINER001")
		mylog.Error(err.Error())
		return "{}"
	}

	// create
	ctx := context.Background()
	cont, err := client.ContainerCreate(ctx, confCont, confHost, &network.NetworkingConfig{}, name)
	if err != nil {
		mylog.Error("ソルバ起動中にエラーが発生しました -> CALLCONTAINER002")
		mylog.Error(err.Error())
		return "{}"
	}

	// start
	err = client.ContainerStart(ctx, cont.ID, types.ContainerStartOptions{})
	if err != nil {
		mylog.Error("ソルバ起動中にエラーが発生しました -> CALLCONTAINER003")
		mylog.Error(err.Error())
		return "{}"
	}
	defer client.ContainerRemove(ctx, cont.ID, types.ContainerRemoveOptions{})

	// wait...
	statusCh, errCh := client.ContainerWait(ctx, cont.ID, container.WaitConditionNotRunning)
	select {
	case <-errCh:
		mylog.Error("ソルバ起動中にエラーが発生しました -> CALLCONTAINER004")
		mylog.Error(err.Error())
		return "{}"
	case <-statusCh:
	}

	// get log
	out, err := client.ContainerLogs(ctx, cont.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		mylog.Error("ソルバ実行中にエラーが発生しました -> CALLCONTAINER005")
		mylog.Error(err.Error())
		return "{}"
	}
	result, _ := ioutil.ReadAll(out)
	return string(result)
}
