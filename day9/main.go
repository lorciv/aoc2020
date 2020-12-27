package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func hasPair(data []int, sum int) bool {
	for i := 0; i < len(data)-1; i++ {
		for j := i + 1; j < len(data); j++ {
			if data[i]+data[j] == sum {
				return true
			}
		}
	}
	return false
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var data []int

	scan := bufio.NewScanner(f)
	// read the preamble
	for i := 0; i < 25; i++ {
		if !scan.Scan() {
			log.Fatal("failed to read preamble")
		}
		text := scan.Text()
		num, _ := strconv.Atoi(text)
		data = append(data, num)
	}

	for scan.Scan() {
		sum, _ := strconv.Atoi(scan.Text())
		if !hasPair(data, sum) {
			fmt.Println(sum)
			break
		}

		data = append(data[1:], sum)
	}

}
