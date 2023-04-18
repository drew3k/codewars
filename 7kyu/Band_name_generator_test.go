package _kyu

import "testing"

func TestBandName(t *testing.T) {
	tests := []struct {
		word string
		exp  string
	}{
		{"knife", "The Knife"},
		{"tart", "Tartart"},
		{"sandles", "Sandlesandles"},
		{"bed", "The Bed"},
	}

	for _, tt := range tests {
		res := bandNameGenerator(tt.word)
		if res != tt.exp {
			t.Errorf("bandNameGenerator (%s) = %s, expected: %s", tt.word, res, tt.exp)
		}
	}

}
