package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const sample = `ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in`

var requiredKeys = []string{
	"byr",
	"iyr",
	"eyr",
	"hgt",
	"hcl",
	"ecl",
	"pid",
}

func Keys(s string) map[string]bool {
	keys := make(map[string]bool)
	for _, f := range strings.Fields(s) {
		keys[strings.Split(f, ":")[0]] = true
	}
	return keys
}

func Validate(keys map[string]bool) bool {
	for _, k := range requiredKeys {
		if !keys[k] {
			return false
		}
	}
	return true
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scan := bufio.NewScanner(f)

	count := 0
	for scan.Scan() {
		line := scan.Text()
		for scan.Scan() {
			nextline := scan.Text()
			if nextline == "" {
				break
			}
			line += " " + nextline
		}

		if Validate(Keys(line)) {
			count++
		}

		line = ""
	}

	fmt.Println(count)
}
