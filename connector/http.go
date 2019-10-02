package connector

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

func httpGet(url string, token string) []byte {
	// setting request
	req, _ := http.NewRequest("GET", url, bytes.NewBuffer([]byte("")))
	req.Header.Set("Authorization", token)

	// http get
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Printf("\x1b[31m[ERROR] HTTP通信(GET)に失敗しました -> HTTPGET001\x1b[0m\n")
		return make([]byte, 0)
	}
	defer res.Body.Close()

	// read data
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("\x1b[31m[ERROR] レスポンスの読み取りに失敗しました -> HTTPGET001\x1b[0m\n")
		return make([]byte, 0)
	}
	return resBody
}

func httpPostJSON(url string, token string, data string) bool {
	// setting request
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(data)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	// http post
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("\x1b[31m[ERROR] HTTP通信(POST)に失敗しました -> HTTPGET001\x1b[0m\n")
		return false
	}
	return resp.StatusCode == 200
}
