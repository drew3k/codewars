package _kyu

import (
	"strconv"
	"strings"
)

func NbDig(n int, d int) (out int) {
	for i := 0; i <= n; i++ {
		out += strings.Count(strconv.Itoa(i*i), strconv.Itoa(d))
	}
	return
}
