package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func isContain(x, pos int) bool {

	posA := []int{pos, pos + 1, pos + 2}

	for _, pos := range posA {
		if pos == x {
			return true
		}
	}
	return false
}

func main() {
	instructions := []int{}

	input, _ := ioutil.ReadFile("input.txt")

	lines := strings.Split(string(input), "\n")

	for _, line := range lines {
		command := strings.Split(line, " ")
		instructions = append(instructions, 0)
		if command[0] == "addx" {
			num, _ := strconv.Atoi(command[1])
			instructions = append(instructions, num)
		}
	}

	// part1
	// sum := 0
	// x := 1
	// for i, inst := range instructions {
	// 	cycle := i + 1
	// 	fmt.Println(cycle, x)
	// 	if cycle%40 == 20 {
	// 		signal := x * cycle
	// 		sum += signal
	// 	}
	// 	x += inst
	// }

	// fmt.Println(sum)

	// part2
	x := 1
	for i, inst := range instructions {
		pos := i % 40
		if isContain(x, pos-1) {
			// fmt.Println("#", x, pos, inst)
			fmt.Print("#")
		} else {
			// fmt.Println(".", x, pos, inst)
			fmt.Print(".")
		}
		x += inst
		if pos == 39 {
			fmt.Println()
		}
	}

}
