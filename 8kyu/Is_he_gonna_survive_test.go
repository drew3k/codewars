package _kyu

import "testing"

func TestHero(t *testing.T) {
	tests := []struct {
		bullets, dragons int
		expected         bool
	}{
		{10, 5, true},
		{7, 4, false},
		{4, 5, false},
		{100, 40, true},
		{1500, 751, false},
		{0, 1, false},
	}

	for _, test := range tests {
		result := Hero(test.bullets, test.dragons)
		if result != test.expected {
			t.Errorf("Hero (%d, %d) = %t; expected %t", test.bullets, test.dragons, result, test.expected)
		}
	}

}
