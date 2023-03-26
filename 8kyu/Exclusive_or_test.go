package _kyu

import "testing"

func TestXor(t *testing.T) {
	tests := []struct {
		a, b     bool
		expected bool
	}{
		{false, false, false},
		{true, false, true},
		{false, true, true},
		{true, true, false},
	}

	for _, tt := range tests {
		res := Xor(tt.a, tt.b)
		if res != tt.expected {
			t.Errorf("Xor (%t, %t) = %t, expected: %t", tt.a, tt.b, res, tt.expected)
		}
	}
}
