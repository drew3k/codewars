package _kyu

import "testing"

func TestFactorial(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{0, 1},
		{1, 1},
		{4, 24},
		{7, 5040},
		{17, 355687428096000},
	}

	for _, tt := range tests {
		res := factorial(tt.n)
		if res != tt.expected {
			t.Errorf("Factorial(%d) = %d, expected: %d", tt.n, res, tt.expected)
		}
	}
}
