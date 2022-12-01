package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func part1(elves []int) {
	// find max
	max := 0
	for _, e := range elves {
		if e > max {
			max = e
		}
	}
	println(max)
}

func part2(elves []int) {
	// find max 3
	sort.Slice(elves, func(i, j int) bool {
		return elves[i] > elves[j]
	})

	fmt.Println(elves[0] + elves[1] + elves[2])
}

func main() {
	elves := make([]int, 0)
	if input, err := ioutil.ReadFile("input.txt"); err == nil {
		lines := strings.Split(string(input), "\n")

		calories := 0
		for _, line := range lines {
			if strings.TrimSpace(line) == "" {
				elves = append(elves, calories)
				calories = 0
			} else {
				trim := strings.TrimSpace(line)
				num, _ := strconv.Atoi(trim)
				calories += num
			}
		}
		elves = append(elves, calories)
	}
	part2(elves)
}
