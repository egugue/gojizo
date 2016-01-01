package api

import (
	"github.com/PuerkitoBio/goquery"
	"net/url"
)

func FetchDictionary(id string) (english, japanese string) {
	url := buildUrl(id)
	doc, _ := goquery.NewDocument(url)

	doc.Find(".NetDicHead").Each(func(_ int, s *goquery.Selection) {
		english = s.Find("span").Text()
	})
	doc.Find(".NetDicBody").Each(func(_ int, s *goquery.Selection) {
		japanese = s.Find("div").Text()
	})

	return english, japanese
}

func buildUrl(id string) string {
	v := url.Values{}
	v.Set("Dic", "EJdict")
	v.Set("Item", id)
	v.Set("Loc", "")
	v.Set("Prof", "XHTML")

	return URL_GET_DIC_ITEM + "?" + v.Encode()
}
