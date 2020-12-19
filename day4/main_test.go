package main

import (
	"strings"
	"testing"
)

func TestKeys(t *testing.T) {
	input := `hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm`
	want := strings.Split("hcl,iyr,eyr,ecl,pid,byr,hgt", ",")
	got := Keys(input)
	for _, k := range want {
		if !got[k] {
			t.Errorf("missing key %q", k)
		}
	}
}
