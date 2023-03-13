package _kyu

import "strconv"

func Strong(n int) string {
	sum := 0
	for _, c := range strconv.Itoa(n) {
		i, _ := strconv.Atoi(string(c))
		sum += Factorial(i)
	}
	if sum == n {
		return "STRONG!!!!"
	}
	return "Not Strong !!"
}

func Factorial(n int) int {
	if n > 0 {
		return n * Factorial(n-1)
	}
	return 1
}
