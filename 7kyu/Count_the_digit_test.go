package _kyu

import "testing"

func TestNbDig(t *testing.T) {
	tests := []struct {
		n, d     int
		expected int
	}{
		{550, 5, 213},
		{5750, 0, 4700},
	}

	for _, tt := range tests {
		res := NbDig(tt.n, tt.d)
		if res != tt.expected {
			t.Errorf("NbDig (%d, %d) = %d, expected: %d", tt.n, tt.d, res, tt.expected)
		}
	}
}
