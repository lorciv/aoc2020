package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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

func ParsePassport(s string) map[string]string {
	pass := make(map[string]string)
	for _, f := range strings.Fields(s) {
		split := strings.Split(f, ":")
		pass[split[0]] = split[1]
	}
	return pass
}

var requiredKeys = []string{
	"byr",
	"iyr",
	"eyr",
	"hgt",
	"hcl",
	"ecl",
	"pid",
}

func Validate(pass map[string]string) bool {
	for _, k := range requiredKeys {
		if pass[k] == "" {
			return false
		}
	}
	return true
}

var regex4digits = regexp.MustCompile("^[0-9]{4}$")
var regex9digits = regexp.MustCompile("^[0-9]{9}$")
var regexHeight = regexp.MustCompile("^[0-9]+(cm|in)$")
var regexColor = regexp.MustCompile("^#[0-9a-f]{6}$")
var eyeColors = map[string]bool{
	"amb": true,
	"blu": true,
	"brn": true,
	"gry": true,
	"grn": true,
	"hzl": true,
	"oth": true,
}

func validateYear(s string, min, max int) bool {
	if !regex4digits.MatchString(s) {
		return false
	}
	n, _ := strconv.Atoi(s)
	return n >= min && n <= max
}

func validateHeight(s string) bool {
	if !regexHeight.MatchString(s) {
		return false
	}

	if strings.HasSuffix(s, "cm") {
		n, _ := strconv.Atoi(strings.TrimSuffix(s, "cm"))
		return n >= 150 && n <= 193
	}
	n, _ := strconv.Atoi(strings.TrimSuffix(s, "in"))
	return n >= 59 && n <= 76
}

func validateColor(s string) bool {
	return regexColor.MatchString(s)
}

func Validate2(pass map[string]string) bool {
	for k, v := range pass {
		ok := false
		switch k {
		case "byr":
			ok = validateYear(v, 1920, 2002)
		case "iyr":
			ok = validateYear(v, 2010, 2020)
		case "eyr":
			ok = validateYear(v, 2020, 2030)
		case "hgt":
			ok = validateHeight(v)
		case "hcl":
			ok = validateColor(v)
		case "ecl":
			ok = eyeColors[v]
		case "pid":
			ok = regex9digits.MatchString(v)
		case "cid":
			ok = true
		}
		if !ok {
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

		pass := ParsePassport(line)
		if Validate(pass) && Validate2(pass) {
			count++
		}

		line = ""
	}

	fmt.Println(count)
}
