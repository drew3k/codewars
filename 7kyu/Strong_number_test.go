package _kyu

import "testing"

func TestStrong(t *testing.T) {
	tests := []struct {
		n        int
		expected string
	}{
		{1, "STRONG!!!!"},
		{2, "STRONG!!!!"},
		{145, "STRONG!!!!"},
		{7, "Not Strong !!"},
		{93, "Not Strong !!"},
		{185, "Not Strong !!"},
	}

	for _, test := range tests {
		result := Strong(test.n)
		if result != test.expected {
			t.Errorf("Strong (%d) = %s, expected: %s", test.n, result, test.expected)
		}
	}
}
