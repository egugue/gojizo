package main

import (
	"encoding/json"
	"github.com/egugue/gojizo/domain/dictionary"
	"goji.io"
	"goji.io/pat"
	"golang.org/x/net/context"
	"net/http"
	"strings"
)

func main() {
	mux := goji.NewMux()
	mux.HandleFuncC(pat.Get("/search/:word"), SearchDictionary)
	http.ListenAndServe(":4000", mux)
}

func SearchDictionary(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	setContentType(w)

	encoder := json.NewEncoder(w)
	word := pat.Param(ctx, "word")
	word = strings.ToLower(word)

	dicList, e := dictionary.FindBy(word)
	if e != nil {
		encoder.Encode(ResError{500, "Internal Server Error"})
		//TODO: status code
		return
	}
	if dicList == nil {
		encoder.Encode(ResError{404, "Not Found"})
		//TODO: status code
		return
	}

	dic, isAttended := dicList.GetBy(word)
	if isAttended {
		encoder.Encode(dic)
		return
	}

	encoder.Encode(dicList)
}

func setContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
}

type ResError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
