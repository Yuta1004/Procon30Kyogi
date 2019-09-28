package connector

import (
	"testing"
)

func TestGetAllBattle(t *testing.T) {
	battleInfo := GetAllBattle("A")
	if battleInfo == nil {
		t.Fail()
	}
}
