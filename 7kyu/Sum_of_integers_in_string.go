package _kyu

import (
	"fmt"
	"regexp"
	"strconv"
)

func SumOfIntegersInString(strng string) int {
	re := regexp.MustCompile("[0-9]+")
	fmt.Println(re.FindAllString(strng, -1))
	sum := 0
	for _, i := range re.FindAllString(strng, -1) {
		number, _ := strconv.Atoi(i)
		sum += number
	}
	return sum
}
