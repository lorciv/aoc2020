package main

import (
	"strings"
	"testing"
)

func TestParseInstr(t *testing.T) {
	tests := []struct {
		input string
		want  Instr
	}{
		{
			"nop +0",
			Instr{"nop", 0},
		},
		{
			"acc +1",
			Instr{"acc", 1},
		},
		{
			"jmp +4",
			Instr{"jmp", 4},
		},
	}

	for _, test := range tests {
		if got := ParseInstr(test.input); got != test.want {
			t.Errorf("ParseInstr(%q) = %v, want %v", test.input, got, test.want)
		}
	}
}

func TestExecute(t *testing.T) {
	prog := `nop +0
	acc +1
	jmp +4
	acc +3
	jmp -3
	acc -99
	acc +1
	jmp -4
	acc +6`

	comp := Computer{}
	comp.Execute(ParseProg(strings.NewReader(prog)))
	if comp.Count != 1 {
		t.Errorf("count %d after execution, want 1", comp.Count)
	}
	if comp.Acc != 5 {
		t.Errorf("acc %d after execution, want 5", comp.Acc)
	}
}
