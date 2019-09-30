package battle

import (
	"github.com/Yuta1004/procon30-kyogi/manager"
)

func copyAllBattleDict() (tmp map[int]manager.Battle) {
	tmp = make(map[int]manager.Battle)
	for key, val := range allBattleDict {
		tmp[key] = val
	}
	return
}
