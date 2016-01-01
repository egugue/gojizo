package util

import (
	"io/ioutil"
	"net/http"
)

func FetchBody(url string) string {
	response, _ := http.Get(url)
	defer response.Body.Close()

	byteArray, _ := ioutil.ReadAll(response.Body)
	return string(byteArray)
}
