package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func parseInstruction1(start int, line string) int {
	chars := strings.Split((strings.TrimSpace(line)), "")
	for _, char := range chars {
		switch char {
		case "U":
			if start > 2 {
				start -= 3
			}
		case "D":
			if start < 7 {
				start += 3
			}
		case "L":
			if start%3 != 1 {
				start -= 1
			}
		case "R":
			if start%3 != 0 {
				start += 1
			}
		}
	}
	return start
}

func part1() {
	start := 5
	if input, err := ioutil.ReadFile("input.txt"); err == nil {
		lines := strings.Split(string(input), "\n")
		for _, line := range lines {
			start = parseInstruction1(start, line)
			fmt.Print(start)
		}
	}
}

func parseInstruction2(start []int, line string) int {
	pad := [5][5]int{{0, 0, 1, 0, 0}, {0, 2, 3, 4, 0}, {5, 6, 7, 8, 9}, {0, 10, 11, 12, 0}, {0, 0, 13, 0, 0}}
	chars := strings.Split((strings.TrimSpace(line)), "")
	for _, char := range chars {
		switch char {
		case "U":
			if start[1] > 0 && pad[start[0]][start[1]-1] != 0 {
				start[1] -= 1
			}
		case "D":
			if start[1] < 4 && pad[start[0]][start[1]+1] != 0 {
				start[1] += 1
			}
		case "L":
			if start[0] > 0 && pad[start[0]-1][start[1]] != 0 {
				start[0] -= 1
			}
		case "R":
			if start[0] < 4 && pad[start[0]+1][start[1]] != 0 {
				start[0] += 1
			}
		}
	}
	return pad[start[1]][start[0]]
}

func part2() {
	start := []int{0, 2}
	if input, err := ioutil.ReadFile("input.txt"); err == nil {
		lines := strings.Split(string(input), "\n")
		for _, line := range lines {
			fmt.Printf("%x", parseInstruction2(start, line))
		}
	}
}

func main() {
	part2()
}
