package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func part1(lines []string) {
	count := 0

	for _, line := range lines {
		sections := strings.Split(line, ",")

		first := strings.Split(sections[0], "-")
		second := strings.Split(sections[1], "-")

		firstStart, _ := strconv.Atoi(first[0])
		firstEnd, _ := strconv.Atoi(first[1])
		secondStart, _ := strconv.Atoi(second[0])
		secondEnd, _ := strconv.Atoi(second[1])
		if firstStart >= secondStart && firstEnd <= secondEnd {
			count += 1
			fmt.Println(line)
			continue
		}
		if firstStart <= secondStart && firstEnd >= secondEnd {
			count += 1
			fmt.Println(line)
			continue
		}
	}
	fmt.Println(count)
}

func part2(lines []string) {
	count := 0

	for _, line := range lines {
		sections := strings.Split(line, ",")

		first := strings.Split(sections[0], "-")
		second := strings.Split(sections[1], "-")

		firstStart, _ := strconv.Atoi(first[0])
		firstEnd, _ := strconv.Atoi(first[1])
		secondStart, _ := strconv.Atoi(second[0])
		secondEnd, _ := strconv.Atoi(second[1])
		if firstStart <= secondEnd && secondStart <= firstEnd {
			count += 1
			fmt.Println(line)
			continue
		}
	}
	fmt.Println(count)
}

func main() {
	start := time.Now()

	input, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(input), "\n")

	// part1(lines)
	part2(lines)

	end := time.Now()
	fmt.Println(end.Sub(start))
}
