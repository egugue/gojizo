package api

import (
	"encoding/xml"
	"net/url"
)

func FetchIdList(word string) []string {
	sdir := searchDicItemResult{}

	url := searchIdListUrl(word)
	xmldoc := fetchBody(url)
	xml.Unmarshal([]byte(xmldoc), &sdir)

	idList := []string{}
	for _, v := range sdir.TitleList.DicItemTitle {
		idList = append(idList, v.ItemID)
	}

	return idList
}

func searchIdListUrl(word string) string {
	v := url.Values{}
	v.Set("Word", word)

	v.Set("Dic", "EJdict")
	v.Set("Scope", "HEADWORD")
	v.Set("Match", "STARTWITH")
	v.Set("Merge", "AND")
	v.Set("Prof", "XHTML")
	v.Set("PageSize", "20")
	v.Set("PageIndex", "0")
	return URL_SEARCH_DIC_ITEM + "?" + v.Encode()
}

type searchDicItemResult struct {
	TotalHitCount int
	ItemCount     int
	TitleList     titleList
}

type titleList struct {
	DicItemTitle []dicItemTitle
}

type dicItemTitle struct {
	ItemID string
}
