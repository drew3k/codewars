package _kyu

import (
	"reflect"
	"testing"
)

func TestMultiple(t *testing.T) {
	tests := []struct {
		ints     []int
		expected []int
	}{
		{[]int{68, -1, 1, -7, 10, 10}, []int{-1, 10}},
		{[]int{22, -6, 32, 82, 9, 25}, []int{-6, 32, 25}},
	}

	for _, test := range tests {
		result := multipleOfIndex(test.ints)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("multipleOfIndex(%v) = %v; expected %v", test.ints, result, test.expected)
		}
	}
}
