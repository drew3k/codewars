package _kyu

import "testing"

func TestSpacify(t *testing.T) {
	tests := []struct {
		s        string
		expected string
	}{
		{"hello world", "h e l l o   w o r l d"},
		{"12345", "1 2 3 4 5"},
		{"Pippi", "P i p p i"},
		{"a", "a"},
		{"", ""},
	}

	for _, tt := range tests {
		res := Spacify(tt.s)
		if res != tt.expected {
			t.Errorf("Spacify (%s) = %s, expected: %s", tt.s, res, tt.expected)
		}
	}
}
