package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Lens struct {
	name    string
	strenth uint8
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

	sequence := strings.Split(data[0], ",")

	part1 := 0
	for _, instruction := range sequence {
		part1 += hash_alg(instruction)
	}
	fmt.Printf("part1: %v\n", part1)

	boxes := make([][]Lens, 256)

	for _, inst := range sequence {
		if string(inst[len(inst)-1]) == "-" {
			// remove lens from box
			lens_str := strings.Replace(inst, "-", "", 1)
			box_no := hash_alg(lens_str)
			// i'm fucking stupid, why did i have > 0 instead of > -1
			if lens_idx := contains(boxes[box_no], lens_str); lens_idx > -1 {
				remove_lens(boxes, box_no, lens_idx)
			}
		} else {
			// add/update lens in box
			lens, err := to_lens(inst)
			if err != nil {
				fmt.Println(err)
			}
			box_no := hash_alg(lens.name)
			if idx := contains(boxes[box_no], lens.name); idx > -1 && idx < 256 {
				boxes[box_no][idx].strenth = lens.strenth
			} else {
				boxes[box_no] = append(boxes[box_no], lens)
			}
		}
	}
	part2 := 0

	for i, b := range boxes {
		if len(b) > 0 {
			for j, l := range b {
				part2 += (i + 1) * (j + 1) * int(l.strenth)
			}
		}
	}
	fmt.Printf("part2: %v\n", part2)

}

func hash_alg(s string) int {
	ascii_codes := []byte(s)
	hash := 0
	for _, chr := range ascii_codes {
		hash = ((hash + int(chr)) * 17) % 256
	}
	return hash
}

func to_lens(s string) (Lens, error) {
	if !strings.Contains(s, "=") {
		return Lens{name: "", strenth: 0}, errors.New(fmt.Sprintf("No '=' found in string '%v'", s))
	}
	splt := strings.Split(s, "=")
	n := splt[0]
	str, _ := strconv.ParseUint(splt[1], 10, 8)
	return Lens{name: n, strenth: uint8(str)}, nil
}

func contains(b []Lens, s string) int {
	for i, l := range b {
		if l.name == s {
			return i
		}
	}
	return -1
}

func remove_lens(b [][]Lens, n int, i int) {
	if i == 0 && len(b) == 1 {
		b[n] = b[n][:0]
	} else if i == 0 {
		b[n] = b[n][1:]
	} else if last := len(b[n]) - 1; i == last {
		b[n] = b[n][:last]
	} else if i > last || i < 0 {
		fmt.Printf("%v out of bounds: %v\n", i, b[n])
	} else {
		b[n] = append(b[n][:i], b[n][i+1:]...)
	}
}

func print_boxes(b [][]Lens) {
	for i, box := range b {
		if len(box) > 0 {
			fmt.Printf("Box %v:", i)
			for _, lens := range box {
				fmt.Printf(" [%v %v]", lens.name, lens.strenth)
			}
			fmt.Printf("\n")
		}
	}
}
