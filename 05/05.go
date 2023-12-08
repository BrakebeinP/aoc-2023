package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

	var seeds []int64
	map_values := make(map[string][]int64)
	title_next := false
	curr_title := ""

	for i, d := range data {
		if i == 0 {
			seeds_text := strings.Split(strings.Split(d, ": ")[1], " ")
			for _, s := range seeds_text {
				n, _ := strconv.ParseInt(s, 10, 64)
				seeds = append(seeds, n)
			}
			continue
		}
		if title_next {
			curr_title := strings.Split(d, " ")[0]
			title_next = false
		}
		if d == "" {
			title_next = true
		}
	}
}
