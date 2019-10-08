package connector

import (
	"github.com/Yuta1004/procon30-kyogi/config"
	"os"
	"os/exec"
	"path"
	"testing"
	"time"
)

var (
	rootPath = path.Join(os.Getenv("GOPATH"), "src/github.com/Yuta1004/procon30-kyogi")
)

func TestConnector(t *testing.T) {
	// setting (config)
	os.Chdir(rootPath)
	token := "procon30_example_token"
	config.SetConfigData(config.Config{
		GameServer: config.GameServer{URL: "http://localhost:8081", Token: token},
	})

	// setting(server)
	cmd := exec.Command("./procon-server_darwin", "--port=8081")
	cmd.Dir = rootPath
	cmd.Start()
	time.Sleep(time.Second)
	defer func() {
		cmd.Process.Kill()
	}()

	// do test
	testGetAllBattle(t, token)
	testGetBattleDetail(t, token)
}

// GetAllBattleのテスト
func testGetAllBattle(t *testing.T, token string) {
	os.Chdir(rootPath)
	battleInfo := *GetAllBattle(token)

	// idx 0
	battle := battleInfo[0]
	if battle.ID != 1 {
		t.Fatal("Error: testAllBattle001")
	}
	if battle.MatchTo != "A高専" {
		t.Fatal("Error: testAllBattle002")
	}

	// idx 1
	battle = battleInfo[1]
	if battle.IntervalMillis != 5000 {
		t.Fatal("Error: testAllBAttle003")
	}
	if battle.TurnMillis != 20000 {
		t.Fatal("Error: testAllBattle004")
	}

	// idx 2
	battle = battleInfo[2]
	if battle.MaxTurn != 60 {
		t.Fatal("Error : testAllBattle005")
	}
}

func testGetBattleDetail(t *testing.T, token string) {
	os.Chdir(rootPath)
	battleDetailInfo := GetBattleDetail(3, token)

	if battleDetailInfo.Height != 10 {
		t.Fatal("Error: testGetBattleDetail001")
	}
	if battleDetailInfo.Width != 10 {
		t.Fatal("Error: testGetBattleDetail002")
	}
	if battleDetailInfo.StartedAtUnixTime != 1561800000 {
		t.Fatal("Error: testGetBattleDetail003")
	}
	if battleDetailInfo.Turn != 2 {
		t.Fatal("Error: testGetBattleDetail004")
	}
}
