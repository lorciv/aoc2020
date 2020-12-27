package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
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

func readData(in io.Reader) []int {
	var data []int
	scan := bufio.NewScanner(in)
	for scan.Scan() {
		num, _ := strconv.Atoi(scan.Text())
		data = append(data, num)
	}
	return data
}

func contiguousSum(data []int, sum int) []int {
	for i := 0; i < len(data)-1; i++ {
		s := data[i]
		for j := i + 1; j < len(data) && s < sum; j++ {
			s += data[j]
			if s == sum {
				return data[i : j+1]
			}
		}
	}
	return nil
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	data := readData(f)
	var sum int
	for i := 25; i < len(data); i++ {
		if !hasPair(data[i-25:i], data[i]) {
			sum = data[i]
			break
		}
	}
	fmt.Println(sum)

	cont := contiguousSum(data, sum)
	sort.Ints(cont)
	fmt.Println(cont)
	fmt.Println(cont[0] + cont[len(cont)-1])

}
