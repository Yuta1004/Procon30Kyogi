package solver

import (
	"os"
	"path"
	"testing"
)

var (
	rootPath = path.Join(os.Getenv("GOPATH"), "src/github.com/Yuta1004/procon30-kyogi")
)

func TestSaveJSON(t *testing.T) {
	os.Chdir(rootPath)
	result := saveJSON("Test", []byte("{\"test\": \"test\"}"))
	if !result {
		t.Fail()
	}
}
