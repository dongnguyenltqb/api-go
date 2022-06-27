package util

import "strings"

func Ref(s string) *string {
	var result = strings.Clone(s)
	return &result
}
