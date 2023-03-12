package _kyu

import "testing"

func TestUnluckyDays(t *testing.T) {
	tests := []struct {
		year     int
		expected int
	}{
		{2015, 3},
		{1986, 1},
	}

	for _, test := range tests {
		result := UnluckyDays(test.year)
		if result != test.expected {
			t.Errorf("UnluckyDays (%d) = %d, expected: %d", test.year, result, test.expected)
		}
	}
}
