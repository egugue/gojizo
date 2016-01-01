package main

import (
	"fmt"
	"github.com/egugue/gojizo/dictionary"
)

func main() {
	word := "adorable"
	//word := "get"
	//word := "adoable"

	dicList := dictionary.FindBy(word)
	dic, isAttended := dicList.GetBy(word)
	if isAttended {
		fmt.Println(dic)
		return
	}

	for _, dic := range dicList {
		fmt.Println(dic)
	}
}
