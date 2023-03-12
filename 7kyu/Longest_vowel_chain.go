package _kyu

func Solve(s string) int {
	n, max := 0, 0
	for _, i := range s {
		switch i {
		case 'a', 'e', 'i', 'o', 'u':
			if n++; n > max {
				max = n
			}
		default:
			n = 0
		}
	}
	return max
}
