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
		next_num := find_pattern(nums, ln == 42)
		part1 += next_num
	}

	fmt.Printf("part1: %v\n", part1)
}

func find_pattern(nums []int64, dbg bool) int64 {
	var diffs []int64
	for i := 0; i < len(nums)-1; i++ {
		diffs = append(diffs, nums[i+1]-nums[i])
	}
	if dbg {
		fmt.Println(diffs)
	}

	var diff_val int64 = 0
	if all_zeroes(diffs) {
		return nums[len(nums)-1]
	} else {
		diff_val = find_pattern(diffs, dbg)
	}

	if dbg {
		fmt.Println(diff_val)
	}
	return nums[len(nums)-1] + diff_val
}

func sum(a []int64) int64 {
	var s int64 = 0
	for _, n := range a {
		s += n
	}
	return s
}

func all_zeroes(a []int64) bool {
	for _, v := range a {
		if v != 0 {
			return false
		}
	}
	return true
}
