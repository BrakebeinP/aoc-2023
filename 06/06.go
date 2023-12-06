package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

	re := regexp.MustCompile(`\d{1,}`)

	raw_race_times := re.FindAllString(data[0], -1)
	raw_distances := re.FindAllString(data[1], -1)

	var race_times []int64
	var distances []int64

	// part1
	var part1_w2w []int
	for i := 0; i < len(raw_race_times); i++ {
		r, _ := strconv.ParseInt(raw_race_times[i], 10, 64)
		d, _ := strconv.ParseInt(raw_distances[i], 10, 64)
		race_times = append(race_times, r)
		distances = append(distances, d)
		ways_to_win := 0
		for attempt := 1; attempt < int(d); attempt++ {
			if attempt*(int(r)-attempt) > int(d) {
				ways_to_win += 1
			}
		}
		part1_w2w = append(part1_w2w, ways_to_win)
	}

	// part 2
	race_time := strings.Join(raw_race_times, "")
	distance := strings.Join(raw_distances, "")

	rt, _ := strconv.ParseInt(race_time, 10, 64)
	ds, _ := strconv.ParseInt(distance, 10, 64)

	part2 := 0

	rs := 0
	for a := 1; a < int(ds); a++ {
		if a*(int(rt)-a) > int(ds) {
			rs = a
			break
		}
	}
	fmt.Printf("rs: %v\n", rs)
	part2 = int(rt) - rs*2 + 1

	part1 := product(part1_w2w)
	fmt.Printf("part1: %v\n", part1)

	fmt.Printf("race_time: %v, distance: %v\n", rt, ds)
	fmt.Printf("part2: %v\n", part2)
}

func product(arr []int) int {
	prod := 0
	for _, p := range arr {
		if prod == 0 {
			prod = p
		} else {
			prod *= p
		}
	}
	return prod
}
