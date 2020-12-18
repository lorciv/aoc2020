package main

import "testing"

func TestFind2(t *testing.T) {
	input := []int{1721, 979, 366, 299, 675, 1456}
	want := 514579
	if got := find2(input); got != want {
		t.Errorf("find2(%v) = %d, want %d", input, got, want)
	}
}

func TestFind3(t *testing.T) {
	input := []int{1721, 979, 366, 299, 675, 1456}
	want := 241861950
	if got := find3(input); got != want {
		t.Errorf("find3(%v) = %d, want %d", input, got, want)
	}
}
