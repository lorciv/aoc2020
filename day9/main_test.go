package main

import "testing"

func TestHasPair(t *testing.T) {
	tests := []struct {
		data []int
		sum  int
		want bool
	}{
		{[]int{35, 20, 15, 25, 47}, 40, true},
		{[]int{20, 15, 25, 47, 40}, 62, true},
		{[]int{15, 25, 47, 40, 62}, 55, true},
		{[]int{95, 102, 117, 150, 182}, 127, false},
	}

	for _, test := range tests {
		if got := hasPair(test.data, test.sum); got != test.want {
			t.Errorf("HasPair(%v, %d) = %v", test.data, test.sum, got)
		}
	}
}
