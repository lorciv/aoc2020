package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const sample = `abc

a
b
c

ab
ac

a
a
a
a

b`

func ParseAnswers(in io.Reader) [][]string {
	var answers [][]string

	scan := bufio.NewScanner(in)
	for scan.Scan() {
		line := scan.Text()
		for scan.Scan() {
			nextline := scan.Text()
			if nextline == "" {
				break
			}
			line += " " + nextline
		}
		answers = append(answers, strings.Fields(line))
	}

	return answers
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	answers := ParseAnswers(f)

	count := 0

	for _, group := range answers {
		n := len(group)
		join := strings.Join(group, "")

		for _, c := range strings.Split(group[0], "") {
			if strings.Count(join, c) >= n {
				count++
			}
		}
	}

	fmt.Println(count)
}
