package dictionary

type DictionaryList []Dictionary

//TODO: bad peformance.
func (dicList DictionaryList) GetBy(word string) (Dictionary, bool) {
	for _, dic := range dicList {
		if dic.English == word {
			return dic, true
		}
	}

	var result Dictionary
	return result, false
}
