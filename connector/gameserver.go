package connector

import (
	"encoding/json"
	"fmt"
	"github.com/Yuta1004/procon30-kyogi/config"
	"os"
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

// GetAllBattle : 自チームが参加している全ての試合情報を取得する
func GetAllBattle(token string) *[]BattleInfo {
	// get data
	config := config.GetConfigData()
	reqURL := config.GameServer.URL + "/matches"
	resBody := httpGet(reqURL, token)

	// json unmarshall
	var battleInfo []BattleInfo
	if err := json.Unmarshal(resBody, &battleInfo); err != nil {
		fmt.Fprintf(os.Stderr, "Could not finished process of unmarshal : %s", err)
		return nil
	}
	return &battleInfo
}

// PostActionData : 行動情報を送信する
func PostActionData(battleID int, token string, actionData string) bool {
	// post data
	config := config.GetConfigData()
	battleIDStr := strconv.Itoa(battleID)
	reqURL := config.GameServer.URL + "/matches/" + battleIDStr + "/action"
	return httpPostJSON(reqURL, token, actionData)
}
