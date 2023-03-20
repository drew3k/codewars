package _kyu

import (
	"fmt"
	"strings"
)

func PartList(l []string) (r string) {
	for i := 1; i < len(l); i++ {
		a := strings.Join(l[0:i], " ")
		b := strings.Join(l[i:], " ")
		r += fmt.Sprintf("(%s, %s)", a, b)
	}
	return
}
