package _kyu

import "strings"

func NoSpace(word string) string {
	return strings.ReplaceAll(word, " ", "")
}
