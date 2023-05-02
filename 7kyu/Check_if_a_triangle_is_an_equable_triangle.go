package _kyu

import "math"

func EquableTriangle(a, b, c int) bool {
	p := a + b + c
	s := p / 2
	area := math.Sqrt(float64(s * (s - a) * (s - b) * (s - c)))
	return p == int(area)
}
