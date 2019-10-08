package battle

import (
	"os"
	"path"
	"testing"
	"time"
)

var (
	rootPath = path.Join(os.Getenv("GOPATH"), "src/github.com/Yuta1004/procon30-kyogi")
)

func TestBManagerExec(t *testing.T) {
	os.Chdir(rootPath)
	ticker := time.NewTicker(10 * time.Second)
	go BManagerExec("A")
	<-ticker.C
}
