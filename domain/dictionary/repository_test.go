package dictionary

import (
	"testing"
)

func TestFindBy(t *testing.T) {
	word := "empty_word"
	actual, _ := FindBy(word)
	if actual != nil {
		t.Errorf("expected : , but actual was %v", actual)
	}

	word = "book"
	actual, _ = FindBy(word)
	if actual == nil {
		t.Errorf("expected is some result, but actual was nil.")
	}
}
