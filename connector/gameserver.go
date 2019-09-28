package connector

import (
	"encoding/json"
	"fmt"
	"github.com/Yuta1004/procon30-kyogi/config"
	"io/ioutil"
	"net/http"
	"os"
)

// BattleInfo : ゲームサーバから受信した試合情報を扱う
type BattleInfo struct {
	ID             int `json:"id"`
	TeamID         int `json:"teamID"`
	TurnMillis     int `json:"turnMillis"`
	IntervalMillis int `json:"intervalMillis"`
	MaxTurn        int `json:"turns"`
	MatchTo        int `json:"matchTo"`
}

// GetAllBattle : 自チームが参加している全ての試合情報を取得する
func GetAllBattle() *[]BattleInfo {
	// URL
	config := config.GetConfigData()
	reqURL := config.GameServer.URL + "/matches"

	// http get
	res, err := http.Get(reqURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not get data from server : %s [%s]\n", err, reqURL)
		return nil
	}
	defer res.Body.Close()

	// read data
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not read data from server : %s\n", err)
		return nil
	}

	// json unmarshall
	var battleInfo []BattleInfo
	if err = json.Unmarshal(resBody, &battleInfo); err != nil {
		fmt.Fprintf(os.Stderr, "Could not finished process of unmarshal : %s", err)
		return nil
	}
	return &battleInfo
}
