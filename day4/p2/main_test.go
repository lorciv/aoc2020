package main

import (
	"strings"
	"testing"
)

func TestValidateYear(t *testing.T) {
	tests := []struct {
		input    string
		min, max int
		want     bool
	}{
		{"1878", 1920, 2002, false},
		{"1950", 1920, 2002, true},
		{"2002", 1920, 2002, true},
		{"2003", 1920, 2002, false},
	}

	for _, test := range tests {
		if got := validateYear(test.input, test.min, test.max); got != test.want {
			t.Errorf("validateBYR(%q, %v, %v) = %v", test.input, test.min, test.max, got)
		}
	}
}

func TestValidateHeight(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"60in", true},
		{"190cm", true},
		{"190in", false},
		{"190", false},
	}

	for _, test := range tests {
		if got := validateHeight(test.input); got != test.want {
			t.Errorf("validateHeight(%q) = %v", test.input, got)
		}
	}
}

func TestValidateColor(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"#123abc", true},
		{"#123abz", false},
		{"123abc", false},
	}

	for _, test := range tests {
		if got := validateColor(test.input); got != test.want {
			t.Errorf("validateColor(%q) = %v", test.input, got)
		}
	}
}

func TestValidate2(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"eyr:1972 cid:100 hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926", false},
		{"iyr:2019 hcl:#602927 eyr:1967 hgt:170cm ecl:grn pid:012533040 byr:1946", false},
		{"hcl:dab227 iyr:2012 ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277", false},
		{"hgt:59cm ecl:zzz eyr:2038 hcl:74454a iyr:2023 pid:3556412378 byr:2007", false},
		{"pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980 hcl:#623a2f", true},
		{"eyr:2029 ecl:blu cid:129 byr:1989 iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm", true},
		{"hcl:#888785 hgt:164cm byr:2001 iyr:2015 cid:88 pid:545766238 ecl:hzl eyr:2022", true},
		{"iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719", true},
	}

	for _, test := range tests {
		pass := ParsePassport(test.input)
		if got := Validate2(pass); got != test.want {
			h := test.input
			if len(h) > 10 {
				h = h[:10] + "..."
			}
			t.Errorf("Validate2(%q) = %v", h, got)
		}
	}
}

func TestParsePassport(t *testing.T) {
	input := `hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm`
	want := strings.Split("hcl,iyr,eyr,ecl,pid,byr,hgt", ",")
	got := ParsePassport(input)
	for _, k := range want {
		if got[k] == "" {
			t.Errorf("missing key %q", k)
		}
	}
}
