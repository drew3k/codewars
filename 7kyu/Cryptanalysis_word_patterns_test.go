package _kyu

import "testing"

func TestWordPattern(t *testing.T) {
	tests := []struct {
		word     string
		expected string
	}{
		{"hello", "0.1.2.2.3"},
		{"heLlo", "0.1.2.2.3"},
		{"helLo", "0.1.2.2.3"},
		{"Hippopotomonstrosesquippedaliophobia", "0.1.2.2.3.2.3.4.3.5.3.6.7.4.8.3.7.9.7.10.11.1.2.2.9.12.13.14.1.3.2.0.3.15.1.13"},
	}

	for _, tt := range tests {
		res := WordPattern(tt.word)
		if res != tt.expected {
			t.Errorf("WordPattern(%s) = %s, expected: %s", tt.word, res, tt.expected)
		}
	}
}
