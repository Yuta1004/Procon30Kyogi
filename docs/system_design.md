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
                                      v
                             +-----------------+
                             |      Docker     |
                             |                 |
                             |     ソルバ実行    |
                             +-----------------+

```

## CUI

- コマンド受付

## BattleManager

- 試合状況を全て管理する
- グローバル変数 `battleList` を使って試合状態を保つ
- システム起動時にgoroutineで呼ばれ、システム終了まで管理を続ける
- チャネルを通してCUI, SolverManagerと通信
- 毎ターン、SovlerManagerに行動情報を依頼する
- 行動情報を受信次第、ServerConnectorに送信を依頼する

```go
var battleList []battleList

type Battle struct {
    ID              int
    TeamID          int
    TurnMillis      int
    IntervalMillis  int
    Turn            int
    MatchTo         string
    SolverCh        chan string
}
```