package _kyu

import "testing"

func TestSumOfIntegersInString(t *testing.T) {
	tests := []struct {
		strng    string
		expected int
	}{
		{"12.4", 16},
		{"h3ll0w0rld", 3},
		{"2 + 3 = ", 5},
		{"Our company made approximately 1 million in gross revenue last quarter.", 1},
		{"The Great Depression lasted from 1929 to 1939.", 3868},
		{"Dogs are our best friends.", 0},
		{"The30quick20brown10f0x1203jumps914ov3r1349the102l4zy dog", 3635},
	}

	for _, test := range tests {
		result := SumOfIntegersInString(test.strng)
		if result != test.expected {
			t.Errorf("SumOfIntegersInString (%s) = %d, expected: %d", test.strng, result, test.expected)
		}
	}
}
