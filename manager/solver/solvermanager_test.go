package solver

import (
	"github.com/Yuta1004/procon30-kyogi/connector"
	"github.com/Yuta1004/procon30-kyogi/manager/battle"
	"testing"
)

func TestExecSolver(t *testing.T) {
	// setup
	agent := connector.Agent{AgentID: 1, X: 0, Y: 0}
	team := connector.Team{TeamID: 1, Agents: []connector.Agent{agent}, AreaPoint: 10, TilePoint: 20}
	action := connector.Action{AgentID: 1, Dx: -1, Dy: -1, Type: "move", Apply: 1, Turn: 1}
	battleDetail := connector.BattleDetailInfo{
		Width:             3,
		Height:            3,
		Turn:              2,
		StartedAtUnixTime: 0,
		Points:            [][]int{{1, 2, 1}, {3, 4, 3}, {1, 2, 1}},
		Tiled:             [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 0}},
		Actions:           []connector.Action{action},
		Teams:             []connector.Team{team},
	}
	battleInfo := connector.BattleInfo{ID: 1, TeamID: 1, TurnMillis: 30000, IntervalMillis: 3000, MaxTurn: 60, MatchTo: "test"}

	// test
	ch := make(chan string, 1)
	go ExecSolver(ch, battle.Battle{Info: &battleInfo, DetailInfo: &battleDetail, Turn: 1})
	<-ch
}
