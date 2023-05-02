package _kyu

import "testing"

func TestEquableTriangle(t *testing.T) {
	tests := []struct {
		a, b, c int
		exp     bool
	}{
		{5, 12, 13, true},
		{2, 3, 4, false},
	}

	for _, tt := range tests {
		res := EquableTriangle(tt.a, tt.b, tt.c)
		if res != tt.exp {
			t.Errorf("EquableTriangle (%d, %d, %d) = %v, expected: %v", tt.a, tt.b, tt.c, res, tt.exp)
		}
	}
}
