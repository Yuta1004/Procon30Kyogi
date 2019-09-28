package connector

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func httpGet(url string) []byte {
	// http get
	res, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not get data from server : %s [%s]\n", err, url)
		return make([]byte, 0)
	}
	defer res.Body.Close()

	// read data
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not read data from server : %s\n", err)
		return make([]byte, 0)
	}
	return resBody
}
