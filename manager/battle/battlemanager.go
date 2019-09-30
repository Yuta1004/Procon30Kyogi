package battle

import (
	"fmt"
	"github.com/Yuta1004/procon30-kyogi/connector"
	"github.com/Yuta1004/procon30-kyogi/manager"
	"github.com/Yuta1004/procon30-kyogi/manager/solver"
	"os"
	"time"
)

var allBattleDict map[int]manager.Battle

// BManagerExec : 名前の通り, 参加している試合全ての管理をする
func BManagerExec(token string) {
	allBattleDict = make(map[int]manager.Battle)
	makeAllBattleDict(token)
	t := time.NewTicker(500 * time.Millisecond)
	for {
		select {
		case <-t.C:
			managerProcess(token)
		}
	}
}

func managerProcess(token string) {
	// main loop
	for _, battle := range copyAllBattleDict() {
		// check battle status
		if battle.Turn <= 0 {
			continue
		}

		// check solver chan
		if battle.SolverCh != nil && len(battle.SolverCh) > 0 {
			result := <-battle.SolverCh
			fmt.Fprintf(os.Stderr, "Solver Output : %s\n", result)
			connector.PostActionData(battle.Info.ID, token, result)
			battle.SolverCh = nil
		}

		// calc elapsedTurn...
		turnMillis := battle.Info.IntervalMillis + battle.Info.TurnMillis
		nowUnix := int(time.Now().UnixNano() / 1000000)
		elapsedTime := nowUnix - battle.DetailInfo.StartedAtUnixTime*1000
		elapsedTurn := int(elapsedTime/turnMillis) + 1

		// update -> exec solver
		if 0 < elapsedTurn && elapsedTurn <= battle.Info.MaxTurn && battle.Turn != elapsedTurn {
			// update battle status
			newerBattle := makeBattleStruct(token, battle.Info.ID)
			newerBattle.Info = battle.Info
			allBattleDict[battle.Info.ID] = newerBattle

			// exec solver
			fmt.Fprintf(os.Stderr, "Exec Solver : %d\n", newerBattle.Info.ID)
			go solver.ExecSolver(newerBattle.SolverCh, newerBattle)
		}
	}
}
