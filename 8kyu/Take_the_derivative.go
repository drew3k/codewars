package _kyu

import "fmt"

func Derive(coefficient, exponent int) string {
	product := coefficient * exponent
	exponent -= 1
	if exponent == 1 {
		return fmt.Sprintf("%dx", product)
	} else {
		return fmt.Sprintf("%dx^%d", product, exponent)
	}
}
