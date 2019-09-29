package solver

import (
	"encoding/json"
	"fmt"
	"github.com/Yuta1004/procon30-kyogi/manager/battle"
	"github.com/docker/docker/client"
	"os"
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
}
