package connector

import (
	"encoding/json"
	"github.com/Yuta1004/procon30-kyogi/config"
	"log"
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
		config := config.GetConfigData()
		reqURL := config.GameServer.URL + "/matches"
		resBody := httpGet(reqURL, token)

		// json unmarshal
		if err := json.Unmarshal(resBody, &battleInfo); err != nil {
			log.Printf("試合情報の取得でエラーが発生しました -> GETALLBATTLE001, Retry: %d\n", retryCnt-1)
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
		config := config.GetConfigData()
		battleIDStr := strconv.Itoa(battleID)
		reqURL := config.GameServer.URL + "/matches/" + battleIDStr
		resBody := httpGet(reqURL, token)

		// json unmarshal
		if err := json.Unmarshal(resBody, &battleDetailInfo); err != nil {
			log.Printf("試合情報の取得でエラーが発生しました -> GETBATTLEDETAIL001, Retry: %d\n", retryCnt-1)
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
		config := config.GetConfigData()
		battleIDStr := strconv.Itoa(battleID)
		reqURL := config.GameServer.URL + "/matches/" + battleIDStr + "/action"
		result = httpPostJSON(reqURL, token, actionData)
		if !result {
			log.Printf("行動情報送信の際にエラーが発生しました -> POSTACTIONDATA001, Retry: %d\n", retryCnt-1)
			continue
		}
		break
	}
	return result
}
