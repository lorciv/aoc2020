package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Grid [][]string

func (g Grid) Hits(right, down int) int {
	count := 0
	row, col := 0, 0
	for row < len(g) {
		if g[row][col%len(g[row])] == "#" {
			count++
		}
		row += down
		col += right
	}
	return count
}

func (g Grid) String() string {
	build := strings.Builder{}
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[i]); j++ {
			build.WriteString(g[i][j])
		}
		build.WriteString("\n")
	}
	return build.String()
}

func ParseGrid(input io.Reader) Grid {
	var grid Grid
	scan := bufio.NewScanner(input)
	for scan.Scan() {
		split := strings.Split(scan.Text(), "")
		grid = append(grid, split)
	}
	return grid
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	grid := ParseGrid(f)

	slopes := []struct {
		right, down int
	}{
		{1, 1},
		{5, 1},
		{3, 1},
		{7, 1},
		{1, 2},
	}

	res := 1
	for _, slope := range slopes {
		res *= grid.Hits(slope.right, slope.down)
	}
	fmt.Println(res)
}
