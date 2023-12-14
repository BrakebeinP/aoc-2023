package main

import (
	"bufio"
	"fmt"
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

	var grid [][]string
	part1 := 0
	for _, d := range data {
		grid = append(grid, strings.Split(d, ""))
	}
	// print_grid(grid)
	// all_north := false
	for {
		moves := 0
		for row := 1; row < len(grid); row++ {
			for col, pos := range grid[row] {
				if pos == "O" && grid[row-1][col] == "." {
					grid[row][col], grid[row-1][col] = grid[row-1][col], grid[row][col]
					moves++
				}
			}
		}
		// print_grid(grid)
		if moves == 0 {
			for i, r := range grid {
				for _, c := range r {
					if c == "O" {
						part1 += (len(grid) - i)
					}
				}
			}
			break
		}

		// if all_north {
		// 	break
		// }
	}
	print_grid(grid)
	fmt.Printf("part1: %v\n", part1)

}

func print_grid(grid [][]string) {
	var p_grid []string
	for _, r := range grid {
		p_grid = append(p_grid, strings.Join(r, " "))
	}
	fmt.Printf("grid:\n%v\n\n", strings.Join(p_grid, "\n"))
}
