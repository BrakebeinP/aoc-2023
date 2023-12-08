package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
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

	dir_counter := 0
	curr_node := nodes["AAA"]
	part1 := 0

	for {
		curr_dir := string(directions[dir_counter])

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

	var part2_nodes []Node
	for k := range nodes {
		if strings.HasSuffix(k, "A") {
			part2_nodes = append(part2_nodes, nodes[k])
		}
	}
	fmt.Println(part2_nodes)

	part2 := 0
	dir_counter = 0
	z_count := 0

	for i := 0; i < 10; i++ {
		curr_dir := string(directions[dir_counter])

		for j, n := range part2_nodes {
			switch curr_dir {
			case "L":
				part2_nodes[j] = nodes[n.left]
			case "R":
				part2_nodes[j] = nodes[n.right]
			}
		}

		part2++

		for _, n := range part2_nodes {
			if strings.HasSuffix(n.name, "Z") {
				z_count++
			}
		}

		if z_count == len(part2_nodes) {
			break
		} else {
			z_count = 0
		}

		dir_counter++
		if !(dir_counter < len(directions)) {
			dir_counter = 0
		}
	}

	fmt.Printf("part1: %v\n", part1)
	fmt.Printf("part2: %v\n", part2)
}
