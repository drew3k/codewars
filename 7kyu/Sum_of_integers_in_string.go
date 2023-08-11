package _kyu

import (
	"regexp"
	"strconv"
)

func SumOfIntegersInString(strng string) int {
	re := regexp.MustCompile("[0-9]+")
	sum := 0
	for _, i := range re.FindAllString(strng, -1) {
		number, _ := strconv.Atoi(i)
		sum += number
	}
	return sum
}
