package solver

import (
	"fmt"
	"os"
	"path"
)

var (
	rootPath = path.Join(os.Getenv("GOPATH"), "src/github.com/Yuta1004/procon30-kyogi")
)

func saveJSON(name string, jsonBody []byte) bool {
	// mkdir
	_ = os.Mkdir(rootPath+"/tmp", 0744)

	// save json
	file, err := os.Create(rootPath + "/tmp/" + name + ".json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not create file : %s", err)
		return false
	}
	file.Write(jsonBody)
	file.Close()
	return true
}
