package jstrings

import (
	"strings"
)

func RemoveLastEnter(x string) string {
	return strings.TrimRight(x, "\n")
}
func RemoveFirstEnter(x string) string {
	return strings.Replace(x, "\n", "", 1)
}
