package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func find2(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if nums[i]+nums[j] == 2020 {
				return nums[i] * nums[j]
			}
		}
	}
	return 0
}

func find3(nums []int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		rem := 2020 - nums[i]
		j, k := i+1, len(nums)-1
		for j < k {
			sum := nums[j] + nums[k]
			if sum == rem {
				return nums[i] * nums[j] * nums[k]
			} else if sum < rem {
				j++
			} else {
				k--
			}
		}
	}
	return 0
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var nums []int
	s := bufio.NewScanner(f)
	for s.Scan() {
		n, err := strconv.Atoi(s.Text())
		if err != nil {
			panic(err)
		}
		nums = append(nums, n)
	}

	res := find2(nums)
	fmt.Println(res)
	fmt.Println(find3(nums))
}
