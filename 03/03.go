package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
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

	var parts_list1 []int64

	for i, r := range data {
		var curr_num strings.Builder
		var num_start int = 0
		for j, c := range r {
			if unicode.IsDigit(c) {
				if curr_num.Len() == 0 {
					num_start = j
				}
				fmt.Fprintf(&curr_num, "%v", c)
			} else if curr_num.Len() != 0 {
				if is_valid_part(curr_num.String(), i, num_start) {
					n, err := strconv.ParseInt(curr_num.String(), 10, 64)
					if err != nil {
						fmt.Println(err)
					}
					parts_list1 = append(parts_list1, n)
				}
				curr_num.Reset()
				num_start = 0
			}
		}
	}
}

func is_valid_part(s string, r int, c int) bool {
	var symbols [10]rune{'#', '$', '%', '&', '@', '-', '+', '*', '/', '='}
}
