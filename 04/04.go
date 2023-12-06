package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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
	var part2 int64 = 0
	card_counts := make([]int64, len(data))
	for i := range card_counts {
		card_counts[i] = 1
	}
	wins_data := make([]int64, len(data))
	for i, r := range data {
		winning_nos := strings.Split(strings.Split(strings.Split(strings.Replace(r, "  ", " ", -1), ": ")[1], " | ")[0], " ")
		player_nos := strings.Split(strings.Split(strings.Split(strings.Replace(r, "  ", " ", -1), ": ")[1], " | ")[1], " ")
		win_cnt := 0
		for _, no := range player_nos {
			for _, w := range winning_nos {
				if no == w {
					win_cnt += 1
					break
				}
			}
		}
		wins_data[i] = int64(win_cnt)
		if win_cnt > 0 {
			part1 += int64(math.Pow(2.0, float64(win_cnt-1)))
			for c := 1; c <= win_cnt; c++ {
				if i+c < len(data) {
					card_counts[i+c] += card_counts[i]
				}
			}
		}
		part2 += card_counts[i]
		// fmt.Printf("%v: %v\n", i, card_counts)
	}
	fmt.Println(wins_data[len(data)-5:])
	fmt.Printf("part1: %v\n", part1)
	fmt.Printf("part2: %v\n", part2)
}
