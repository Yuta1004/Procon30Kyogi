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
	for _, battle := range allBattleDict {
		// check solver chan
		if battle.SolverCh != nil && len(battle.SolverCh) > 0 {
			result := <-battle.SolverCh
			fmt.Fprintf(os.Stderr, "Solver Output : %s\n", result)
			connector.PostActionData(battle.Info.ID, token, result)
			battle.SolverCh = nil
		}

		// check to change of turn
		turnMillis := battle.Info.IntervalMillis + battle.Info.TurnMillis
		nowUnix := int(time.Now().UnixNano() / 1000000)
		elapsedTime := nowUnix - battle.DetailInfo.StartedAtUnixTime*1000
		elapsedTurn := int(elapsedTime/turnMillis) + 1
		if 0 < elapsedTurn && elapsedTurn <= battle.Info.MaxTurn && battle.Turn != elapsedTurn {
			// update battle status
			newerBattle := newBattle(token, battle.Info.ID)
			newerBattle.Info = battle.Info
			allBattleDict[battle.Info.ID] = newerBattle

			// exec solver
			fmt.Fprintf(os.Stderr, "Exec Solver : %d\n", newerBattle.Info.ID)
			go solver.ExecSolver(newerBattle.SolverCh, newerBattle)
		}
	}
}

func makeAllBattleDict(token string) {
	battleInfoList := connector.GetAllBattle(token)
	for _, battleInfo := range *battleInfoList {
		battle := newBattle(token, battleInfo.ID)
		battle.Info = &battleInfo
		allBattleDict[battleInfo.ID] = battle
	}
}

func newBattle(token string, battleID int) manager.Battle {
	fmt.Fprintf(os.Stderr, "Get Data : %d\n", battleID)
	battleDetailInfo := connector.GetBattleDetail(battleID, token)
	return manager.Battle{
		DetailInfo: battleDetailInfo,
		Turn:       battleDetailInfo.Turn,
		SolverCh:   make(chan string, 10),
	}
}
