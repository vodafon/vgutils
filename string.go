package vgutils

import (
	"crypto/rand"
	"encoding/hex"
	"strings"
)

func RandomHEXString(size int) string {
	dst := make([]byte, size)
	rand.Read(dst)
	return hex.EncodeToString(dst)
}

func StringSliceContains(sl []string, e string) bool {
	for _, v := range sl {
		if strings.Contains(v, e) {
			return true
		}
	}
	return false
}

func StringContainsSlice(sl []string, e string) bool {
	for _, v := range sl {
		if strings.Contains(e, v) {
			return true
		}
	}
	return false
}

func StringSliceInclude(sl []string, e string) bool {
	for _, v := range sl {
		if v == e {
			return true
		}
	}
	return false
}

func StringSliceCompact(sl []string) []string {
	res := []string{}
	for _, v := range sl {
		if strings.TrimSpace(v) != "" {
			res = append(res, v)
		}
	}
	return res
}
