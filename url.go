package vgutils

import (
	"net/url"
)

func AddParamsToURL(u, params string) (string, error) {
	uu, err := url.Parse(u)
	if err != nil {
		return "", err
	}
	if uu.Path == "" {
		uu.Path = "/"
	}
	query, err := url.ParseQuery(params)
	if err != nil {
		return "", err
	}
	orig := uu.Query()
	for k, vals := range query {
		for _, v := range vals {
			orig.Add(k, v)
		}
	}
	uu.RawQuery = orig.Encode()
	return uu.String(), nil
}
