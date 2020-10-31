package vgutils

import (
	"testing"
)

type URLTT struct {
	u   string
	p   string
	exp string
}

func TestAddParamsToURL(t *testing.T) {
	params := "key1=value1&key2=value2"
	tt := []URLTT{
		{"https://example.com", params, "https://example.com/?key1=value1&key2=value2"},
		{"https://example.com/?abc=cba", params, "https://example.com/?abc=cba&key1=value1&key2=value2"},
	}
	for i, v := range tt {
		res, err := AddParamsToURL(v.u, v.p)
		if err != nil {
			t.Fatal(err)
		}
		if res != v.exp {
			t.Errorf("%d: expected %q, got %q", i, v.exp, res)
		}
	}
}
