package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Policy struct {
	char     string
	min, max int
}

func (p Policy) Validate(psw string) bool {
	n := strings.Count(psw, p.char)
	return n >= p.min && n <= p.max
}

func (p Policy) Validate2(psw string) bool {
	count := 0
	for _, pos := range []int{p.min - 1, p.max - 1} {
		if psw[pos:pos+1] == p.char {
			count++
		}
	}
	return count == 1
}

func ParsePolicy(s string) Policy {
	split := strings.Split(s, " ")
	rng := strings.Split(split[0], "-")
	min, _ := strconv.Atoi(rng[0])
	max, _ := strconv.Atoi(rng[1])
	return Policy{
		char: split[1],
		min:  min, max: max,
	}
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	count := 0
	count2 := 0
	input := bufio.NewScanner(f)
	for input.Scan() {
		split := strings.Split(input.Text(), ":")
		pol := ParsePolicy(split[0])
		psw := strings.TrimSpace(split[1])

		if pol.Validate(psw) {
			count++
		}
		if pol.Validate2(psw) {
			count2++
		}
	}

	fmt.Println("policy 1:", count)
	fmt.Println("policy 2:", count2)
}
