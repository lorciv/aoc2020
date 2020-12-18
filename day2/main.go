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

func ParsePolicy(s string) Policy {
	parts := strings.Split(s, " ")
	rng := strings.Split(parts[0], "-")
	min, _ := strconv.Atoi(rng[0])
	max, _ := strconv.Atoi(rng[1])
	return Policy{
		char: parts[1],
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

	input := bufio.NewScanner(f)
	for input.Scan() {
		parts := strings.Split(input.Text(), ":")
		pol := ParsePolicy(parts[0])
		psw := strings.TrimSpace(parts[1])

		if pol.Validate(psw) {
			count++
		}
	}

	fmt.Println(count)
}
