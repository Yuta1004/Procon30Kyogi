package connector

import (
	"encoding/json"
	"github.com/Yuta1004/procon30-kyogi/config"
	"github.com/Yuta1004/procon30-kyogi/mylog"
	"strconv"
)

// BattleInfo : ゲームサーバから受信した試合情報を扱う
type BattleInfo struct {
	ID             int    `json:"id"`
	TeamID         int    `json:"teamID"`
	TurnMillis     int    `json:"turnMillis"`
	IntervalMillis int    `json:"intervalMillis"`
	MaxTurn        int    `json:"turns"`
	MatchTo        string `json:"matchTo"`
}

// BattleDetailInfo : ゲームサーバから受信した試合情報詳細を扱う
type BattleDetailInfo struct {
	Width             int      `json:"width"`
	Height            int      `json:"height"`
	Turn              int      `json:"turn"`
	StartedAtUnixTime int      `json:"startedAtUnixTime"`
	Points            [][]int  `json:"points"`
	Tiled             [][]int  `json:"tiled"`
	Actions           []Action `json:"actions"`
	Teams             []Team   `json:"teams"`
}

// Action : 行動情報を扱う
type Action struct {
	AgentID int    `json:"agentID"`
	Dx      int    `json:"dx"`
	Dy      int    `json:"dy"`
	Type    string `json:"type"`
	Apply   int    `json:"apply"`
	Turn    int    `json:"turn"`
}

// Team : チーム情報を扱う
type Team struct {
	TeamID    int     `json:"teamID"`
	Agents    []Agent `json:"agents"`
	AreaPoint int     `json:"areaPoint"`
	TilePoint int     `json:"tilePoint"`
}

// Agent : エージェント情報を扱う
type Agent struct {
	AgentID int `json:"agentID"`
	X       int `json:"x"`
	Y       int `json:"y"`
}

// GetAllBattle : 自チームが参加している全ての試合情報を取得する
func GetAllBattle(token string) *[]BattleInfo {
	var battleInfo []BattleInfo
	for retryCnt := 3; retryCnt > 0; retryCnt-- {
		// get data
		conf := config.GetConfigData()
		reqURL := conf.GameServer.URL + "/matches"
		resBody := httpGet(reqURL, token)

		// json unmarshal
		if err := json.Unmarshal(resBody, &battleInfo); err != nil {
			continue
		}
		break
	}
	return &battleInfo
}

// GetBattleDetail : 試合情報詳細を取得する
func GetBattleDetail(battleID int, token string) BattleDetailInfo {
	var battleDetailInfo BattleDetailInfo
	for retryCnt := 3; retryCnt > 0; retryCnt-- {
		// get data
		conf := config.GetConfigData()
		battleIDStr := strconv.Itoa(battleID)
		reqURL := conf.GameServer.URL + "/matches/" + battleIDStr
		resBody := httpGet(reqURL, token)

		// json unmarshal
		if err := json.Unmarshal(resBody, &battleDetailInfo); err != nil {
			continue
		}
		break
	}
	return battleDetailInfo
}

// PostActionData : 行動情報を送信する
func PostActionData(battleID int, token string, actionData string) bool {
	result := false
	for retryCnt := 3; retryCnt > 0; retryCnt-- {
		conf := config.GetConfigData()
		battleIDStr := strconv.Itoa(battleID)
		reqURL := conf.GameServer.URL + "/matches/" + battleIDStr + "/action"
		result = httpPostJSON(reqURL, token, actionData)
		if !result {
			continue
		}
		break
	}

	if result {
		mylog.Info("行動情報送信が正常に完了しました -> Token: %s, BattleID: %d", token, battleID)
	} else {
		mylog.Error("行動情報送信に失敗しました -> POSTACTIONDATA001")
	}
	return result
}

// CheckToken : トークンが正しいか問い合わせる
func CheckToken(token string) {
	conf := config.GetConfigData()
	url := conf.GameServer.URL + "/ping"
	res := httpGet(url, token)
	mylog.Notify("トークン検証結果: %s", res)
}
