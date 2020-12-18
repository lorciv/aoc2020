package main

import (
	"strings"
	"testing"
)

const sample = `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`

func TestHits(t *testing.T) {
	grid := ParseGrid(strings.NewReader(sample))

	tests := []struct {
		right, down int
		want        int
	}{
		{1, 1, 2},
		{3, 1, 7},
		{5, 1, 3},
		{7, 1, 4},
		{1, 2, 2},
	}

	for _, test := range tests {
		if got := grid.Hits(test.right, test.down); got != test.want {
			t.Errorf("Hits(%d, %d) = %d, want %d", test.right, test.down, got, test.want)
		}
	}
}
