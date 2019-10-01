package battle

import (
	"github.com/Yuta1004/procon30-kyogi/connector"
	"github.com/Yuta1004/procon30-kyogi/manager"
	"github.com/Yuta1004/procon30-kyogi/manager/solver"
	"log"
	"time"
)

var allBattleDict map[int]manager.Battle

// BManagerExec : 名前の通り, 参加している試合全ての管理をする
func BManagerExec(token string) {
	log.Printf("[INFO] BattleManager起動...\n")
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
		// check solver chan
		if battle.SolverCh != nil && len(battle.SolverCh) > 0 {
			result := <-battle.SolverCh
			log.Printf("[INFO] ソルバの実行が終了しました -> BattleID: %d\n", battle.Info.ID)
			go connector.PostActionData(battle.Info.ID, token, result)
			battle.SolverCh = nil
		}

		// calc elapsedTurn...
		turnMillis := battle.Info.IntervalMillis + battle.Info.TurnMillis
		nowUnix := int(time.Now().UnixNano() / 1000000)
		elapsedTime := nowUnix - battle.DetailInfo.StartedAtUnixTime*1000
		elapsedTurn := int(elapsedTime/turnMillis) + 1

		// update -> exec solver
		if 0 <= elapsedTime && 1 <= elapsedTurn && elapsedTurn <= battle.Info.MaxTurn && battle.Turn != elapsedTurn {
			newerBattle := makeBattleStruct(token, battle.Info.ID)
			if newerBattle.Turn != battle.Turn {
				// update battle status
				newerBattle.Info = battle.Info
				newerBattle.ProcessErrCnt = 0
				allBattleDict[battle.Info.ID] = newerBattle
				log.Printf("\x1b[32m[NOTIFY] 次ターンに移行しました -> BattleID: %d, Turn : %d\x1b[0m\n", newerBattle.Info.ID, newerBattle.Turn)

				// exec solver
				go solver.ExecSolver(newerBattle.SolverCh, newerBattle)
			}
		}

		// relief failed...
		if battle.ProcessErrCnt >= 5 {
			log.Printf("\x1b[31m[ERROR] 試合情報の復旧に失敗しました. 該当試合の更新を中断します -> BattleID: %d\n", battle.Info.ID)
			delete(allBattleDict, battle.Info.ID)
			continue
		}

		// relief
		if battle.DetailInfo.StartedAtUnixTime == 0 {
			log.Printf("\x1b[31m[ERROR] 試合情報の復旧を行います -> BattleID: %d\n", battle.Info.ID)
			newerBattle := makeBattleStruct(token, battle.Info.ID)
			newerBattle.Info = battle.Info
			newerBattle.ProcessErrCnt++
			allBattleDict[battle.Info.ID] = newerBattle
		}
	}
}
