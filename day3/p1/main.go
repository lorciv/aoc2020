package main

import (
	"bufio"
	"fmt"
	"os"
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

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	count := 0
	pos := 0

	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		hit := line[pos:pos+1] == "#"
		if hit {
			count++
		}
		pos += 3
		pos = pos % len(line)
	}

	fmt.Println(count)
}
