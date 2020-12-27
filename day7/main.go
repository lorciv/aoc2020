package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const sample = `light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`

type Rule struct {
	Container string
	Content   string
	Quantity  int
}

var rules []Rule

func (r Rule) String() string {
	return fmt.Sprintf("%q -> %d %q", r.Container, r.Quantity, r.Content)
}

func directContainers(bag string) []string {
	var content []string
	for _, r := range rules {
		if r.Content == bag {
			content = append(content, r.Container)
		}
	}
	return content
}

func allContainers(targ string) []string {
	var visited = make(map[string]bool)
	var queue = directContainers(targ)

	for len(queue) > 0 {
		visited[queue[0]] = true
		queue = append(queue, directContainers(queue[0])...)
		queue = queue[1:]
	}

	keys := make([]string, 0, len(visited))
	for k := range visited {
		keys = append(keys, k)
	}
	return keys
}

// ParseRules parses the rules contained in a line.
// input: "light red bags contain 1 bright white bag, 2 muted yellow bags."
func ParseRules(str string) error {
	str = strings.TrimSuffix(str, ".")

	split := strings.Split(str, " bags contain ")
	container := split[0]
	if split[1] == "no other bags" {
		return nil
	}

	contents := strings.Split(split[1], ", ")
	for _, c := range contents {
		// "1 bright white bag"
		rule := Rule{
			Container: container,
		}

		sub := strings.SplitN(c, " ", 2)
		rule.Quantity, _ = strconv.Atoi(sub[0])
		if rule.Quantity == 1 {
			sub[1] = strings.TrimSuffix(sub[1], " bag")
		} else {
			sub[1] = strings.TrimSuffix(sub[1], " bags")
		}
		rule.Content = sub[1]

		rules = append(rules, rule)
	}

	return nil
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scan := bufio.NewScanner(f)
	for scan.Scan() {
		ParseRules(scan.Text())
	}

	for _, r := range rules {
		fmt.Println(r)
	}

	fmt.Println("direct containers", directContainers("shiny gold"))
	all := allContainers("shiny gold")
	fmt.Println("all containers", all)
	fmt.Println(len(all))
}
