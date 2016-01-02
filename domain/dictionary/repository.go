package dictionary

import (
	"github.com/egugue/gojizo/infra/api"
)

func FindBy(word string) (DictionaryList, error) {
	idList, e := api.FetchIdList(word)

	if e != nil {
		return nil, e
	}
	if len(idList) < 1 {
		return nil, nil
	}

	var english, japanese string
	var dicList DictionaryList
	for _, id := range idList {
		english, japanese = api.FetchDictionary(id)
		dicList = append(dicList, Dictionary{id, english, japanese})
	}

	return dicList, nil
}
