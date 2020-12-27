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

func (c *Computer) Execute(prog []Instr) error {
	visited := make(map[int]bool)

	for !visited[c.Count] {
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
			return fmt.Errorf("unknown operation %q", instr.Op)
		}

		fmt.Printf("%v | count = %d, acc = %d\n", instr, c.Count, c.Acc)
	}

	return nil
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	prog := ParseProg(f)
	comp := Computer{}
	comp.Execute(prog)
}
