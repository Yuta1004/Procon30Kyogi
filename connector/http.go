package connector

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func httpGet(url string, token string) []byte {
	// setting request
	req, _ := http.NewRequest("GET", url, bytes.NewBuffer([]byte("")))
	req.Header.Set("Authorization", token)

	// http get
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s, %s\n", url, err)
		return make([]byte, 0)
	}
	defer res.Body.Close()

	// read data
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
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
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
	return resp.StatusCode == 200
}
