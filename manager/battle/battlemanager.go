package battle

import (
	"encoding/json"
	"github.com/Yuta1004/procon30-kyogi/connector"
	"github.com/Yuta1004/procon30-kyogi/manager"
	"github.com/Yuta1004/procon30-kyogi/manager/solver"
	"github.com/Yuta1004/procon30-kyogi/mylog"
	"time"
)

var allBattleDict map[int]manager.Battle

// BManagerExec : 名前の通り, 参加している試合全ての管理をする
func BManagerExec(token string) {
	// panic handling
	defer func() {
		if err := recover(); err != nil {
			mylog.Error("回復不可能なエラー(panic)が発生しました! BattleManagerを終了します!")
			mylog.Error("%s", err)
		}
	}()

	// setting...
	mylog.Info("BattleManager起動...")
	allBattleDict = make(map[int]manager.Battle)
	MakeAllBattleDict(token)

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

		// update -> exec solver -> relief
		if checkNeedUpdateBattle(battle) {
			newerBattle := makeBattleStruct(token, battle.Info.ID)
			if newerBattle.Turn != battle.Turn {
				newerBattle.Info = battle.Info
				newerBattle.ProcessErrCnt = 0
				allBattleDict[battle.Info.ID] = newerBattle
				outBattleLog(newerBattle)
				go solver.ExecSolver(newerBattle.SolverCh, newerBattle)
			}
		}
		reliefBattle(token, battle)
	}
}

func checkSolver(battle manager.Battle) string {
	// receive data
	solverRes := <-battle.SolverCh
	var tmp interface{}

	// valid json
	if err := json.Unmarshal([]byte(solverRes), &tmp); err != nil {
		mylog.Error("ソルバが正常に終了しませんでした -> BattleID: %d", battle.Info.ID)
		mylog.Error(solverRes)
		solverRes = ""
	} else {
		mylog.Info("ソルバの実行が正常に終了しました -> BattleID: %d", battle.Info.ID)
	}
	return solverRes
}

func reliefBattle(token string, battle manager.Battle) {
	// relief failed...
	if battle.ProcessErrCnt >= 5 {
		mylog.Error("試合情報の復旧に失敗しました. 該当試合の更新を中断します -> BattleID: %d", battle.Info.ID)
		delete(allBattleDict, battle.Info.ID)
		return
	}

	// relief
	if battle.DetailInfo.StartedAtUnixTime == 0 {
		mylog.Error("試合情報の復旧を行います -> BattleID: %d, ErrCnt: %d", battle.Info.ID, battle.ProcessErrCnt+1)
		newerBattle := makeBattleStruct(token, battle.Info.ID)
		newerBattle.Info = battle.Info
		newerBattle.ProcessErrCnt = battle.ProcessErrCnt + 1
		allBattleDict[battle.Info.ID] = newerBattle
	}
}
