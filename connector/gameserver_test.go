package connector

import (
	"os"
	"path"
	"testing"
)

var (
	rootPath = path.Join(os.Getenv("GOPATH"), "src/github.com/Yuta1004/procon30-kyogi")
)

func TestGetAllBattle(t *testing.T) {
	os.Chdir(rootPath)
	battleInfo := GetAllBattle("A")
	if battleInfo == nil {
		t.Fail()
	}
}
