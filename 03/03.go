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
}

func main() {
	f, err := os.Open("test.txt")

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

	var parts_list1 []Number

	part1 := 0

	for i, d := range data {
		nums_in_row := re.FindAllString(d, -1)
		num_idx := re.FindAllIndex([]byte(d), -1)
		// fmt.Printf("r%v:%v\n", i, nums_in_row)

		if nums_in_row != nil {
			for j, num := range nums_in_row {
				value, _ := strconv.ParseInt(num, 10, 0)
				curr_num := Number{num: num, val: int(value), row: i, start: num_idx[j][0], end: num_idx[j][1]}
				if is_valid_part(curr_num, data) {
					parts_list1 = append(parts_list1, curr_num)
					part1 += curr_num.val
				}

				// fmt.Println(curr_num)
			}
		}
	}

	fmt.Printf("part1: %v\n", part1)
}

func is_valid_part(n Number, d []string) bool {
	// symbols := []string{"#", "$", "%", "&", "@", "-", "+", "*", "/", "="}
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
