package _kyu

import (
	s "strings"
	"unicode"
)

func solve(str string) string {
	lower := 0
	upper := 0
	for _, i := range str {
		if unicode.IsLower(i) {
			lower += 1
		} else {
			upper += 1
		}
	}
	if upper > lower {
		return s.ToUpper(str)
	}
	return s.ToLower(str)
}
