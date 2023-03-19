package _kyu

import (
	"fmt"
	"strings"
)

func WordPattern(word string) string {
	m := map[rune]int{}
	pos := 0
	d := ""
	var sb strings.Builder
	for _, r := range strings.ToLower(word) {
		if _, ok := m[r]; !ok {
			m[r] = pos
			pos++
		}
		_, err := fmt.Fprintf(&sb, "%s%d", d, m[r])
		if err != nil {
			return ""
		}
		d = "."
	}
	return sb.String()
}
