package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Instr struct {
	Op  string
	Arg int
}

func ParseInstr(s string) Instr {
	fields := strings.Fields(s)
	arg, _ := strconv.Atoi(fields[1])
	return Instr{
		Op:  fields[0],
		Arg: arg,
	}
}

func ParseProg(in io.Reader) []Instr {
	scan := bufio.NewScanner(in)
	prog := make([]Instr, 0)
	for scan.Scan() {
		prog = append(prog, ParseInstr(scan.Text()))
	}
	return prog
}

type Computer struct {
	Count int
	Acc   int
}

func (c *Computer) Execute(prog []Instr) (bool, error) {
	visited := make(map[int]bool)

	for c.Count < len(prog) {
		if visited[c.Count] {
			return true, nil
		}
		visited[c.Count] = true

		instr := prog[c.Count]
		switch instr.Op {
		case "acc":
			c.Acc += instr.Arg
			c.Count++
		case "jmp":
			c.Count += instr.Arg
		case "nop":
			c.Count++
		default:
			return false, fmt.Errorf("unknown operation %q", instr.Op)
		}
	}

	return false, nil
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	prog := ParseProg(f)

	for i := 0; i < len(prog); i++ {
		for prog[i].Op == "acc" {
			i++
		}
		prog2 := make([]Instr, len(prog))
		copy(prog2, prog)
		if prog2[i].Op == "nop" {
			prog2[i].Op = "jmp"
		} else {
			prog2[i].Op = "nop"
		}
		comp := Computer{}
		loop, _ := comp.Execute(prog2)
		if loop {
			fmt.Println("loop!")
		} else {
			fmt.Println(comp)
			break
		}
	}

}
