package solver

import (
	"testing"
)

func TestSaveJSON(t *testing.T) {
	result := SaveJSON("Test", []byte("{\"test\": \"test\"}"))
	if !result {
		t.Fail()
	}
}
