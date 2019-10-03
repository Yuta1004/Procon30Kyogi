package battle

import (
	"encoding/json"
	"github.com/Yuta1004/procon30-kyogi/connector"
	"github.com/Yuta1004/procon30-kyogi/manager"
	"github.com/Yuta1004/procon30-kyogi/manager/solver"
	"log"
	"time"
)

var allBattleDict map[int]manager.Battle

// BManagerExec : 名前の通り, 参加している試合全ての管理をする
func BManagerExec(token string) {
	// panic handling
	defer func() {
		if err := recover(); err != nil {
			log.Printf("\x1b[31m[ERROR] 回復不可能なエラー(panic)が発生しました! BattleManagerを終了します!\x1b[0m\n")
			log.Printf("\x1b[31m%s\x1b[0m\n", err)
		}
	}()

	// setting...
	log.Printf("[INFO] BattleManager起動...\n")
	allBattleDict = make(map[int]manager.Battle)
	makeAllBattleDict(token)

	// mainloop
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
			solverRes := checkSolver(battle)
			go connector.PostActionData(battle.Info.ID, token, solverRes)
			battle.SolverCh = nil
		}

		// update -> exec solver
		elapsedTime, elapsedTurn := calcTimeStatus(battle)
		if 0 <= elapsedTime && 1 <= elapsedTurn && elapsedTurn <= battle.Info.MaxTurn && battle.Turn != elapsedTurn {
			newerBattle := makeBattleStruct(token, battle.Info.ID)
			if newerBattle.Turn != battle.Turn {
				// update battle status
				newerBattle.Info = battle.Info
				newerBattle.ProcessErrCnt = 0
				allBattleDict[battle.Info.ID] = newerBattle

				// output log
				outBattleLog(newerBattle)

				// exec solver
				go solver.ExecSolver(newerBattle.SolverCh, newerBattle)
			}
		}

		// relief
		reliefBattle(token, battle)
	}
}

func checkSolver(battle manager.Battle) string {
	// receive data
	solverRes := <-battle.SolverCh
	var tmp []interface{}

	// valid json
	if err := json.Unmarshal([]byte(solverRes), &tmp); err != nil {
		log.Printf("\x1b[31m[ERROR] ソルバが正常に終了しませんでした -> BattleID: %d\n", battle.Info.ID)
		return "{}"
	}
	log.Printf("[INFO] ソルバの実行が正常に終了しました -> BattleID: %d\n", battle.Info.ID)
	return solverRes
}

func outBattleLog(battle manager.Battle) {
	score := getScore(battle)
	log.Printf("\x1b[32m[NOTIFY] 次ターンに移行しました -> BattleID: %d, Turn : %d\x1b[0m\n", battle.Info.ID, battle.Turn)
	log.Printf(
		"[INFO] 試合情報 -> \x1b[1mBattleID: %d, \x1b[31m自チーム: %d (A %d, T %d), \x1b[34m相手チーム: %d (A %d, T %d)\x1b[0m\n",
		battle.Info.ID, score[0][0]+score[0][1], score[0][0], score[0][1], score[1][0]+score[1][1], score[1][0], score[1][1],
	)
}

func reliefBattle(token string, battle manager.Battle) {
	// relief failed...
	if battle.ProcessErrCnt >= 5 {
		log.Printf("\x1b[31m[ERROR] 試合情報の復旧に失敗しました. 該当試合の更新を中断します -> BattleID: %d\x1b[0m\n", battle.Info.ID)
		delete(allBattleDict, battle.Info.ID)
		return
	}

	// relief
	if battle.DetailInfo.StartedAtUnixTime == 0 {
		log.Printf("\x1b[31m[ERROR] 試合情報の復旧を行います -> BattleID: %d, ErrCnt: %d\x1b[0m\n", battle.Info.ID, battle.ProcessErrCnt+1)
		newerBattle := makeBattleStruct(token, battle.Info.ID)
		newerBattle.Info = battle.Info
		newerBattle.ProcessErrCnt = battle.ProcessErrCnt + 1
		allBattleDict[battle.Info.ID] = newerBattle
	}
}
