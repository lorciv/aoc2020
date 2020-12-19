package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// 111
// 110
// 101
// 100
// 011
// 010
// 001
// 000

type Seat [10]string

func ParseSeat(s string) (Seat, error) {
	var seat Seat
	split := strings.Split(s, "")
	if len(split) != 10 {
		return Seat{}, fmt.Errorf("cannot parse %q: malformed seat", s)
	}
	copy(seat[:], split)
	return seat, nil
}

func (s Seat) Row() int {
	row := 0
	for i := 0; i < 7; i++ {
		row = row << 1
		if s[i] == "B" {
			row++
		}
	}
	return row
}

func (s Seat) Col() int {
	col := 0
	for i := 7; i < 10; i++ {
		col = col << 1
		if s[i] == "R" {
			col++
		}
	}
	return col
}

func (s Seat) ID() int {
	return s.Row()*8 + s.Col()
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	maxID := 0

	scan := bufio.NewScanner(f)
	for scan.Scan() {
		seat, err := ParseSeat(scan.Text())
		if err != nil {
			log.Fatal(err)
		}
		id := seat.ID()
		if id > maxID {
			maxID = id
		}
	}

	fmt.Println(maxID)
}
