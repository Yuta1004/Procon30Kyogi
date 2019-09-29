# SystemDesign

## システム全体図

```

+-----------------+          +-----------------+          +--------------------+
|       CUI       |          |  BattleManager  |          |  ServerConnector   |
|                 | <----->  |                 | <----->  |                    |
|    コマンド受付   |          |     試合管理     |          | ゲームサーバと通信する |
+-----------------+          +-----------------+          +-------------------+
                                      ^
                                      |
                                      v
                             +-----------------+
                             |  SolverManager  |
                             |                 |
                             |    ソルバ管理     |
                             +-----------------+
                                      ^
                                      |
                                [Docker API]
                                      |
                                      v
                             +-----------------+
                             |      Docker     |
                             |                 |
                             |     ソルバ実行    |
                             +-----------------+

```

## CUI

- コマンド受付
- システム起動時に呼ばれる
- チャネルを生成し、BattleManagerをgoroutineで起動する

## BattleManager

- 試合状況を全て管理する
- グローバル変数 `battleList` を使って試合状態を保つ
- システム起動時にgoroutineで呼ばれ、システム終了まで管理を続ける
- チャネルを通してCUI, SolverManagerと通信
- 毎ターン、SovlerManagerに行動情報を依頼する
- 行動情報を受信次第、ServerConnectorに送信を依頼する

```go
var battleList []battleList

// Battle : 試合情報を扱う
type Battle struct {
    Info            BattleInfo
    Turn            int
    SolverCh        chan string
}
```

## SolverManager

- BattleManagerにgoroutineで呼ばれる
- 試合状況を判断し、適切なリソース割り当てを行う (優先度低)
- Solver実行終了後はチャネルを通してBattleManagerへ結果を返す
- Solver実行時間の管理も行う

## ServerConnector

- ゲームサーバと通信する
- ゲームデータ受信, 行動情報送信を行う

```go
// BattleInfo : ゲームサーバから受信した試合情報を扱う
type BattleInfo struct {
	ID             int
	TeamID         int
	TurnMillis     int
	IntervalMillis int
	MaxTurn        int
	MatchTo        string
}

// BattleDetailInfo : ゲームサーバから受信した試合情報詳細を扱う
type BattleDetailInfo struct {
	Width             int
	Height            int
	Turn              int
	StartedAtUnixTime int
	Points            [][]int
	Tiled             [][]int
	Actions           []Action
	Teams             []Team
}

// Action : 行動情報を扱う
type Action struct {
	AgentID int
	Dx      int
	Dy      int
	Type    string
	Apply   int
	Turn    int
}

// Team : チーム情報を扱う
type Team struct {
	TeamID    int
	Agents    []Agent
	AreaPoint int
	TilePoint int
}

// Agent : エージェント情報を扱う
type Agent struct {
	AgentID int
	X       int
	Y       int
}
```