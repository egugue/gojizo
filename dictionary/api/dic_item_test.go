package api

import (
	"testing"
)

//TODO : test
func testBuildUrl(t *testing.T) {
	gi := "givenItem"
	actual := buildUrl(gi)
	expected := URL_GET_DIC_ITEM + "?Dic=EJdict&Item=" + gi + "&Loc=&Prof=XHTML"

	if actual != expected {
		t.Errorf("\nexpected : %v\n actual : %v\n", expected, actual)
	}
}
