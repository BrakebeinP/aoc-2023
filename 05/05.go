package main

import (
	"bufio"
	"fmt"
	"os"
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

	map_values := make(map[string][]string)

	for i, d := range data {
		if i == 0 {

		}
	}
}
