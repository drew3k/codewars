package _kyu

import "testing"

func TestPotatoes(t *testing.T) {
	tests := []struct {
		p0, w0, p1 int
		expected   int
	}{
		{99, 100, 98, 50},
		{82, 127, 80, 114},
	}

	for _, tt := range tests {
		result := Potatoes(tt.p0, tt.w0, tt.p1)
		if result != tt.expected {
			t.Errorf("Potatoes (%d, %d, %d) = %d, expected: %d", tt.p0, tt.w0, tt.p1, result, tt.expected)
		}
	}
}
