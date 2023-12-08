package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type Node struct {
	name  string
	left  string
	right string
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

	re := regexp.MustCompile(`\w{3}`)

	var directions string
	nodes := make(map[string]Node)

	for i, d := range data {
		if i == 0 {
			directions = d
			continue
		} else if i == 1 {
			continue
		}
		n := re.FindAllString(d, -1)
		// fmt.Println(n)
		nodes[n[0]] = Node{name: n[0], left: n[1], right: n[2]}
	}

	fmt.Println(directions, len(directions))

	dir_counter := 0
	curr_node := nodes["AAA"]
	part1 := 0
	for {
		curr_dir := string(directions[dir_counter])
		// fmt.Printf("node %v:\n  left: %v\n  right: %v\n chosen dir: %v\n", curr_node.name, curr_node.left, curr_node.right, curr_dir)

		switch curr_dir {
		case "L":
			curr_node = nodes[curr_node.left]
		case "R":
			curr_node = nodes[curr_node.right]
		}

		part1++

		if curr_node.name == "ZZZ" {
			break
		}

		dir_counter++
		if !(dir_counter < len(directions)) {
			dir_counter = 0
		}
	}

	fmt.Printf("part1: %v", part1)
}
