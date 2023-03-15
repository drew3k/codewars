package _kyu

import "testing"

func TestHeron(t *testing.T) {
	tests := []struct {
		a, b, c  float64
		expected float64
	}{
		{3.0, 4.0, 5.0, 6.0},
	}

	for _, tt := range tests {
		res := Heron(tt.a, tt.b, tt.c)
		if res != tt.expected {
			t.Errorf("Heron (%f, %f, %f) = %f, expected: %f", tt.a, tt.b, tt.c, res, tt.expected)
		}
	}

}
