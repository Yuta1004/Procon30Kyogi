package connector

import (
	"testing"
)

func TestGetAllBattle(t *testing.T) {
	battleInfo := GetAllBattle()
	if battleInfo == nil {
		t.Fail()
	}
}
