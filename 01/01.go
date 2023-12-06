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

	var part1 int64 = 0
	var part2 int64 = 0
	for _, r := range data {
		val := get_num(r)
		part1 += val

		r2 := part2conv(r)
		val2 := get_num(r2)
		part2 += val2
	}

	fmt.Printf("part1: %d\n", part1)
	fmt.Printf("part2: %d\n", part2)
}

func part2conv(s string) string {
	var numbers = [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	convert_map := map[string]string{"one": "on1e", "two": "tw2o", "three": "thre3e", "four": "fo4our", "five": "fi5ve", "six": "si6x", "seven": "se7ven", "eight": "ei8ght", "nine": "ni9ne"}
	new_s := s
	for i := 0; i < len(new_s); i++ {
		for _, n := range numbers {
			l := len(n)

			if i+l <= len(new_s) {
				if contains(numbers, new_s[i:i+l]) {
					new_s = strings.Replace(new_s, new_s[i:i+l], convert_map[new_s[i:i+l]], 1)
				}
			}
		}
	}
	return new_s
}

func contains(s [9]string, t string) bool {
	for _, a := range s {
		if a == t {
			return true
		}
	}
	return false
}

func get_num(s string) int64 {
	num := ""
	last_num := ""
	for _, i := range s {
		if unicode.IsDigit(i) {
			if num == "" {
				num = string(i)
				last_num = string(i)
			} else {
				last_num = string(i)
			}
		}
	}
	num += last_num
	val, err := strconv.ParseInt(num, 10, 64)
	if err != nil {
		fmt.Printf("error parsing %v", s)
	}
	return val
}
