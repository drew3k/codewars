package _kyu

func Spacify(s string) string {
	space := ""
	for i, char := range s {
		if i > 0 {
			space += " "
		}
		space += string(char)
	}
	return space
}
