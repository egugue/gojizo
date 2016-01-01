package dictionary

import (
	"github.com/egugue/gojizo/dictionary/api"
)

func FindBy(word string) DictionaryList {
	var dicList []Dictionary

	idList := api.FetchIdList(word)

	if len(idList) < 1 {
		return dicList
	}

	var english, japanese string
	for _, id := range idList {
		english, japanese = api.FetchDictionary(id)
		dicList = append(dicList, Dictionary{id, english, japanese})
	}

	return dicList
}
