package api

import (
	"io/ioutil"
	"net/http"
)

func fetchBody(url string) string {
	response, _ := http.Get(url)
	defer response.Body.Close()

	byteArray, _ := ioutil.ReadAll(response.Body)
	return string(byteArray)
}
