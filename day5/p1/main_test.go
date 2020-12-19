package main

import "testing"

func TestSeat(t *testing.T) {
	tests := []struct {
		input        string
		row, col, id int
	}{
		{"FBFBBFFRLR", 44, 5, 357},
		{"BFFFBBFRRR", 70, 7, 567},
		{"FFFBBBFRRR", 14, 7, 119},
		{"BBFFBBFRLL", 102, 4, 820},
	}

	for _, test := range tests {
		seat, err := ParseSeat(test.input)
		if err != nil {
			t.Fatal("could not load location")
		}
		if got := seat.Row(); got != test.row {
			t.Errorf("%q.Row() = %d, want %d", test.input, got, test.row)
		}
		if got := seat.Col(); got != test.col {
			t.Errorf("%q.Col() = %d, want %d", test.input, got, test.col)
		}
		if got := seat.ID(); got != test.id {
			t.Errorf("%q.ID() = %d, want %d", test.input, got, test.id)
		}
	}
}
