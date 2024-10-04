package tests

import (
	"firstproject/internal/lcm"
	"testing"
)

func TestFindLCM(t *testing.T) {
	tests := []struct {
		a, b, c  int
		expected int
	}{
		{5, 7, 15, 105},
		{100, 50, 1, 100},
		{3, 9, 27, 27},
	}

	for _, test := range tests {
		result := lcm.FindLCM(test.a, test.b, test.c)
		if result != test.expected {
			t.Errorf("FindLCM(%d, %d, %d) = %d; want %d", test.a, test.b, test.c, result, test.expected)
		}
	}
}
