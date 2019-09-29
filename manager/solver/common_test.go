package solver

import (
	"testing"
)

func TestSaveJSON(t *testing.T) {
	result := saveJSON("Test", []byte("{\"test\": \"test\"}"))
	if !result {
		t.Fail()
	}
}
