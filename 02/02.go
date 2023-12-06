package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(os.Args)
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

	part1 := 0
	var part2 int64 = 0
	for i, v := range data {
		game_choices := strings.Split(strings.Split(v, ": ")[1], "; ")
		// fmt.Printf("%v: %v\n", i+1, game_choices)
		if part1possible(game_choices) {
			// fmt.Println(i + 1)
			part1 += i + 1
		}
		part2 += part2power(game_choices)
	}

	fmt.Printf("part1: %v\n", part1)
	fmt.Printf("part2: %v\n", part2)
}

func part1possible(s []string) bool {
	for _, rnd := range s {
		colors := strings.Split(rnd, ", ")
		for _, clr := range colors {
			cnt := strings.Split(clr, " ")
			n, err := strconv.ParseInt(cnt[0], 10, 64)
			if err != nil {
				fmt.Println(err)
			}
			if (cnt[1] == "red" && n > 12) || (cnt[1] == "green" && n > 13) || (cnt[1] == "blue" && n > 14) {
				return false
			}
		}
	}
	return true
}

func part2power(s []string) int64 {
	var red_min int64 = 0
	var green_min int64 = 0
	var blue_min int64 = 0

	for _, rnd := range s {
		colors := strings.Split(rnd, ", ")
		for _, clr := range colors {
			cnt := strings.Split(clr, " ")
			n, err := strconv.ParseInt(cnt[0], 10, 64)
			if err != nil {
				fmt.Println(err)
			}
			if cnt[1] == "red" && red_min < n {
				red_min = n
			} else if cnt[1] == "green" && green_min < n {
				green_min = n
			} else if cnt[1] == "blue" && blue_min < n {
				blue_min = n
			}
		}
	}
	// fmt.Printf(" r: %v\n g: %v\n b: %v\n", red_min, green_min, blue_min)

	return red_min * blue_min * green_min
}
