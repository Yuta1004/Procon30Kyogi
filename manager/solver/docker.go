package solver

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"io/ioutil"
	"os"
)

func callContainer(confCont *container.Config, confHost *container.HostConfig, name string) string {
	// crate client
	client, err := client.NewClientWithOpts(client.WithVersion("1.40"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return "Error"
	}

	// create
	ctx := context.Background()
	cont, err := client.ContainerCreate(ctx, confCont, confHost, &network.NetworkingConfig{}, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return "Error"
	}

	// start
	err = client.ContainerStart(ctx, cont.ID, types.ContainerStartOptions{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return "Error"
	}
	defer client.ContainerRemove(ctx, cont.ID, types.ContainerRemoveOptions{})

	// wait...
	statusCh, errCh := client.ContainerWait(ctx, cont.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return "Error"
	case <-statusCh:
	}

	// get log
	out, err := client.ContainerLogs(ctx, cont.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return "Error"
	}
	result, _ := ioutil.ReadAll(out)
	return string(result)
}
