package _kyu

import "testing"

func TestLongestVowelChain(t *testing.T) {
	tests := []struct {
		s        string
		expected int
	}{
		{"codewarriors", 2},
		{"suoidea", 3},
		{"ultrarevolutionariees", 3},
		{"strengthlessnesses", 1},
		{"cuboideonavicuare", 2},
		{"chrononhotonthuooaos", 5},
		{"iiihoovaeaaaoougjyaw", 8},
	}

	for _, test := range tests {
		result := Solve(test.s)
		if result != test.expected {
			t.Errorf("Solve (%s) = %d, expected %d", test.s, result, test.expected)
		}
	}
}
