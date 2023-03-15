package _kyu

import "testing"

func TestSolve(t *testing.T) {
	tests := []struct {
		str      string
		expected string
	}{
		{"coDe", "code"},
		{"CODe", "CODE"},
		{"coDE", "code"},
		{"a", "a"},
		{"Z", "Z"},
		{"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz"},
	}

	for _, tt := range tests {
		result := solve(tt.str)
		if result != tt.expected {
			t.Errorf("solve (%s) = %s, expected: %s", tt.str, result, tt.expected)
		}
	}
}
