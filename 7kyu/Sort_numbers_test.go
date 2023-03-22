package _kyu

import (
	"reflect"
	"testing"
)

func TestSortNumbers(t *testing.T) {
	tests := []struct {
		numbers  []int
		expected []int
	}{
		{[]int{1, 2, 10, 50, 5}, []int{1, 2, 5, 10, 50}},
		{[]int{}, []int{}},
	}

	for _, tt := range tests {
		res := SortNumbers(tt.numbers)
		if !reflect.DeepEqual(tt.expected, tt.numbers) {
			t.Errorf("SortNumbers (%v) = %v, expected: %v", tt.numbers, res, tt.expected)
		}
	}
}
