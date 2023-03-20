package _kyu

import (
	"reflect"
	"testing"
)

func TestPartList(t *testing.T) {
	tests := []struct {
		l        []string
		expected string
	}{
		{[]string{"I", "wish", "I", "hadn't", "come"}, "(I, wish I hadn't come)(I wish, I hadn't come)(I wish I, hadn't come)(I wish I hadn't, come)"},
		{[]string{"cdIw", "tzIy", "xDu", "rThG"}, "(cdIw, tzIy xDu rThG)(cdIw tzIy, xDu rThG)(cdIw tzIy xDu, rThG)"},
	}

	for _, tt := range tests {
		res := PartList(tt.l)
		if !reflect.DeepEqual(res, tt.expected) {
			t.Errorf("PartList (%v) = %s, expected: %s", tt.l, res, tt.expected)
		}
	}
}
