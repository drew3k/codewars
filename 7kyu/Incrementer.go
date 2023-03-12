package _kyu

func Incrementer(n []int) []int {
	res := []int{}
	for i, num := range n {
		sum := num + (i + 1)
		res = append(res, sum%10)
	}
	return res
}
