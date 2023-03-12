package _kyu

import (
	"reflect"
	"testing"
)

func TestIncrementer(t *testing.T) {
	tests := []struct {
		n        []int
		expected []int
	}{
		{[]int{1, 2, 3}, []int{2, 4, 6}},
		{[]int{4, 6, 7, 1, 3}, []int{5, 8, 0, 5, 8}},
		{[]int{4, 6, 7, 1, 3}, []int{5, 8, 0, 5, 8}},
		{[]int{3, 6, 9, 8, 9}, []int{4, 8, 2, 2, 4}},
		{[]int{}, []int{}},
	}

	for _, test := range tests {
		result := Incrementer(test.n)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Incrementer (%v) = %v, expected %v", test.n, result, test.expected)
		}
	}
}
