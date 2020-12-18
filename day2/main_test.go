package main

import "testing"

func TestValidate2(t *testing.T) {
	tests := []struct {
		pol  Policy
		psw  string
		want bool
	}{
		{Policy{"a", 1, 3}, "abcde", true},
		{Policy{"b", 1, 3}, "cdefg", false},
		{Policy{"c", 2, 9}, "ccccccccc", false},
	}

	for _, test := range tests {
		if got := test.pol.Validate2(test.psw); got != test.want {
			t.Errorf("Validate2(%v, %s) = %v, want %v", test.pol, test.psw, got, test.want)
		}
	}
}

func TestValidate(t *testing.T) {
	tests := []struct {
		pol  Policy
		psw  string
		want bool
	}{
		{Policy{"a", 1, 3}, "abcde", true},
		{Policy{"b", 1, 3}, "cdefg", false},
		{Policy{"c", 2, 9}, "ccccccccc", true},
	}

	for _, test := range tests {
		if got := test.pol.Validate(test.psw); got != test.want {
			t.Errorf("Validate(%v, %s) = %v, want %v", test.pol, test.psw, got, test.want)
		}
	}
}

func TestParsePolicy(t *testing.T) {
	tests := []struct {
		input string
		want  Policy
	}{
		{"1-3 a", Policy{"a", 1, 3}},
		{"1-3 b", Policy{"b", 1, 3}},
		{"2-9 c", Policy{"c", 2, 9}},
	}

	for _, test := range tests {
		if got := ParsePolicy(test.input); got != test.want {
			t.Errorf("ParsePolicy(%q) = %v, want %v", test.input, got, test.want)
		}
	}
}
