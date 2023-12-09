package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	var data []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		curr_row := scanner.Text()
		if err != nil {
			fmt.Println(err)
		}
		data = append(data, curr_row)
	}

	var part1 int64 = 0

	for ln, d := range data {
		var nums []int64
		textnums := strings.Split(d, " ")
		for _, n := range textnums {
			num, _ := strconv.ParseInt(n, 10, 64)
			nums = append(nums, num)
		}
		next_num := find_pattern(nums)
		part1 += next_num
		fmt.Printf("ln %v: %v\n", ln+1, next_num)
	}

	fmt.Printf("part1: %v\n", part1)
}

func find_pattern(nums []int64) int64 {
	var diffs []int64
	for i := 0; i < len(nums)-1; i++ {
		diffs = append(diffs, nums[i+1]-nums[i])
	}

	nums_val := 0
	var return_val int64 = 0
	for _, diff := range diffs {
		nums_val += int(diff)
	}
	if nums_val == 0 {
		return nums[len(nums)-1]
	} else {
		return_val = find_pattern(diffs)
	}
	return nums[len(nums)-1] + return_val
}
