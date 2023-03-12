package _kyu

import "testing"

func TestContamination(t *testing.T) {
	tests := []struct {
		text, char string
		expected   string
	}{
		{"abc", "z", "zzz"},
		{"", "z", ""},
		{"abc", "", ""},
		{"_3ebzgh4", "&", "&&&&&&&&"},
		{"//case", " ", "      "},
	}

	for _, test := range tests {
		result := Contamination(test.text, test.char)
		if result != test.expected {
			t.Errorf("Contamination (%s, %s) = %s, expected %s", test.text, test.char, result, test.expected)
		}
	}
}
