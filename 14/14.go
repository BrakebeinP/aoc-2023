package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Direction int

const (
	North = iota
	West
	South
	East
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

	var grid_orig [][]string
	var gridp1 [][]string
	var gridp2 [][]string
	for _, d := range data {
		grid_orig = append(grid_orig, strings.Split(d, ""))
		gridp1 = append(gridp1, strings.Split(d, ""))
		gridp2 = append(gridp2, strings.Split(d, ""))
	}

	for {
		moves := 0

		shift_north(gridp1)

		if moves == 0 {
			break
		}
	}

	fmt.Printf("part1: %v\n", calc_score(gridp1))

	cycle_count := 1000000000
	directions := []Direction{North, West, South, East}

	gridmap := make(map[string]map[Direction]string)

	curr_grid := grid_to_string(gridp2)
	first_rotation := 0
	var next_dir Direction
	for cnt := 1; cnt <= cycle_count; cnt++ {
		for d := 0; d < len(directions); d++ {
			dir := directions[d]
			if _, ok := gridmap[curr_grid]; ok {
				if grd, ok2 := gridmap[curr_grid][dir]; ok2 {
					curr_grid = grd
					if first_rotation == 0 {
						first_rotation = cnt
						next_dir = dir
						break
					}
				} else {
					gridmap[curr_grid][dir] = shift_grid(gridp2, dir)
					curr_grid = gridmap[curr_grid][dir]
				}
			} else {
				gridmap[curr_grid] = make(map[Direction]string)
				gridmap[curr_grid][dir] = shift_grid(gridp2, dir)
				curr_grid = gridmap[curr_grid][dir]
			}
		}
		if first_rotation != 0 {
			break
		}

		if cnt%1000000 == 0 {
			fmt.Println(cnt)
		}

	}
	fmt.Println(first_rotation)
	// print_grid(gridp2)
	rem := cycle_count % first_rotation

	for i := 0; i < rem; i++ {
		for d := 0; d < len(directions); d++ {
			dir := directions[d]
			shift_grid(gridp2, dir)
		}
	}

	fmt.Printf("part2: %v\n", calc_score(gridp2))
}

func print_grid(grid [][]string) {
	var p_grid []string
	for _, r := range grid {
		p_grid = append(p_grid, strings.Join(r, " "))
	}
	fmt.Printf("grid:\n%v\n\n", strings.Join(p_grid, "\n"))
}

func calc_score(grid [][]string) int {
	score := 0
	for i, r := range grid {
		for _, c := range r {
			if c == "O" {
				score += (len(grid) - i)
			}
		}
	}
	return score
}

func swap_pos(r1 int, c1 int, r2 int, c2 int, grid [][]string) {
	grid[r1][c1], grid[r2][c2] = grid[r2][c2], grid[r1][c1]
}

func is_equal(s1, s2 [][]string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if strings.Join(s1[i], "") != strings.Join(s2[i], "") || len(s1[i]) != len(s2[i]) {
			return false
		}
	}

	return true
}

func shift_north(grid [][]string) {
	for col := 0; col < len(grid[0]); col++ {
		for row := 1; row < len(grid); row++ {
			if grid[row][col] == "O" {
				best_row := 0
				for swap_row := row; swap_row > 0; swap_row-- {
					if grid[swap_row-1][col] == "#" || grid[swap_row-1][col] == "O" {
						best_row = swap_row
						break
					} else if swap_row == 0 && grid[0][col] == "." {
						best_row = 0
					}
				}
				if best_row != row {
					swap_pos(best_row, col, row, col, grid)
				}
			}
		}
	}
}

func shift_west(grid [][]string) {
	for row := 0; row < len(grid); row++ {
		for col := 1; col < len(grid[row]); col++ {
			if grid[row][col] == "O" {
				best_col := 0
				for swap_col := col; swap_col > 0; swap_col-- {
					if grid[row][swap_col-1] == "#" || grid[row][swap_col-1] == "O" {
						best_col = swap_col
						break
					} else if swap_col == 0 && grid[row][0] == "." {
						best_col = 0
					}
				}
				if best_col != col {
					swap_pos(row, col, row, best_col, grid)
				}
			}
		}
	}
}

func shift_south(grid [][]string) {
	for col := 0; col < len(grid[0]); col++ {
		for row := len(grid) - 1; row > -1; row-- {
			if grid[row][col] == "O" {
				best_row := len(grid) - 1
				for swap_row := row; swap_row < len(grid)-1; swap_row++ {
					if grid[swap_row+1][col] == "#" || grid[swap_row+1][col] == "O" {
						best_row = swap_row
						break
					} else if l := len(grid) - 1; swap_row == l && grid[l][col] == "." {
						best_row = l
					}
				}
				if best_row != row {
					swap_pos(best_row, col, row, col, grid)
				}
			}
		}
	}
}

func shift_east(grid [][]string) {
	for row := 0; row < len(grid); row++ {
		for col := len(grid[row]) - 1; col > -1; col-- {
			if grid[row][col] == "O" {
				best_col := len(grid[row]) - 1
				for swap_col := col; swap_col < len(grid[row])-1; swap_col++ {
					if grid[row][swap_col+1] == "#" || grid[row][swap_col+1] == "O" {
						best_col = swap_col
						break
					} else if l := len(grid[row]) - 1; swap_col == l && grid[row][l] == "." {
						best_col = l
					}
				}
				if best_col != col {
					swap_pos(row, col, row, best_col, grid)
				}
			}
		}
	}
}

func shift_grid(grid [][]string, dir Direction) string {
	switch dir {
	case North:
		shift_north(grid)
	case West:
		shift_west(grid)
	case South:
		shift_south(grid)
	case East:
		shift_east(grid)
	}

	return grid_to_string(grid)
}

func grid_to_string(grid [][]string) string {
	var str []string
	for _, r := range grid {
		str = append(str, strings.Join(r, ""))
	}
	return strings.Join(str, "\n")
}
