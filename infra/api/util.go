package api

import (
	"io/ioutil"
	"net/http"
)

func fetchBody(url string) (string, error) {
	resp, e := http.Get(url)
	if e != nil {
		return "", e
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	return string(byteArray), nil
}
