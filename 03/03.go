package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

type Number struct {
	num   string
	val   int
	row   int
	start int
	end   int
	star  string
}

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

	re := regexp.MustCompile(`\d{1,}`)
	re2 := regexp.MustCompile(`\*`)

	var parts_list []Number

	part1 := 0

	for i, d := range data {
		nums_in_row := re.FindAllString(d, -1)
		num_idx := re.FindAllIndex([]byte(d), -1)

		// part1
		if nums_in_row != nil {
			for j, num := range nums_in_row {
				value, _ := strconv.ParseInt(num, 10, 0)
				curr_num := Number{num: num, val: int(value), row: i, start: num_idx[j][0], end: num_idx[j][1], star: ""}
				if is_valid_part(curr_num, data) {
					parts_list = append(parts_list, curr_num)
					part1 += curr_num.val
				}
			}
		}
	}

	// part2
	part2 := 0
	for r, d2 := range data {
		star_idx := re2.FindAllIndex([]byte(d2), -1)
		if star_idx == nil {
			continue
		}
		for _, star := range star_idx {
			var nums []int
			c := star[0]

			r_offsets := []int{-1, 0, 1}
			c_offsets := []int{-1, 0, 1}

			for _, rr := range r_offsets {
				for _, cc := range c_offsets {
					if unicode.IsDigit(rune(data[r+rr][c+cc])) {
						for _, part := range parts_list {
							if part.row == r+rr && part.start <= c+cc && c+cc < part.end {
								if (len(nums) > 0 && nums[len(nums)-1] != part.val) || len(nums) == 0 {
									nums = append(nums, part.val)
								}
							}
						}
					}
				}
			}
			if len(nums) != 2 {
				continue
			}
			part2 += nums[0] * nums[1]
		}
	}

	fmt.Printf("part1: %v\n", part1)
	fmt.Printf("part2: %v\n", part2)
}

func is_valid_part(n Number, d []string) bool {
	var start int
	var end int
	if n.start-1 < 0 {
		start = 0
	} else {
		start = n.start - 1
	}
	if n.end == len(d[n.row]) {
		end = n.end - 1
	} else {
		end = n.end
	}
	rr := []int{-1, 0, 1}
	for _, i := range rr {
		r := n.row + i
		if r >= 0 && r < len(d) {
			for c := start; c <= end; c++ {
				if string(d[r][c]) != "." && !unicode.IsDigit(rune(d[r][c])) {
					return true
				}
			}
		}
	}
	return false
}
