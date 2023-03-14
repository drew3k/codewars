package _kyu

import "testing"

func TestRepeatStr(t *testing.T) {
	tests := []struct {
		repetitions int
		value       string
		expected    string
	}{
		{4, "a", "aaaa"},
		{3, "hello ", "hello hello hello "},
		{2, "abc", "abcabc"},
	}

	for _, test := range tests {
		result := RepeatStr(test.repetitions, test.value)
		if result != test.expected {
			t.Errorf("RepeatStr (%d, %s) = %s, expected: %s", test.repetitions, test.value, result, test.expected)
		}
	}
}
