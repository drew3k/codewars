package _kyu

import (
	"testing"
)

func TestDerive(t *testing.T) {
	tests := []struct {
		coefficient, exponent int
		expected              string
	}{
		{2, 3, "6x^2"},
		{5, 1, "5x^0"},
		{4, 0, "0x^-1"},
	}

	for _, test := range tests {
		result := Derive(test.coefficient, test.exponent)
		if result != test.expected {
			t.Errorf("Derive(%d, %d) = %s; expected %s", test.coefficient, test.exponent, result, test.expected)
		}
	}
}
