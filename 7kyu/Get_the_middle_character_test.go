package _kyu

import "testing"

func TestGetMiddle(t *testing.T) {
	tests := []struct {
		s   string
		exp string
	}{
		{"test", "es"},
		{"testing", "t"},
		{"middle", "dd"},
		{"A", "A"},
	}

	for _, tt := range tests {
		res := GetMiddle(tt.s)
		if res != tt.exp {
			t.Errorf("GetMiddle (%s) = %s, expected: %s", tt.s, res, tt.exp)
		}
	}
}
