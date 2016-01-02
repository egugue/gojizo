package api

import (
	"testing"
)

func TestFetchBody(t *testing.T) {
	_, e := fetchBody("unknown_host_xxxx.com/foo/bar")
	if e == nil {
		t.Errorf("exptected is error, but it was success.")
	}
}
