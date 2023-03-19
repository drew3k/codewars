package _kyu

import "testing"

func TestNoSpace(t *testing.T) {
	tests := []struct {
		word     string
		expected string
	}{
		{"8 j 8   mBliB8g  imjB8B8  jl  B", "8j8mBliB8gimjB8B8jlB"},
		{"8j aam", "8jaam"},
		{"8aaaaa dddd r     ", "8aaaaaddddr"},
		{"8 8 Bi fk8h B 8 BB8B B B  B888 c hl8 BhB fd", "88Bifk8hB8BB8BBBB888chl8BhBfd"},
	}

	for _, tt := range tests {
		res := NoSpace(tt.word)
		if res != tt.expected {
			t.Errorf("NoSpace (%s) = %s, expected: %s", tt.word, res, tt.expected)
		}
	}
}
