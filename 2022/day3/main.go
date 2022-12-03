package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

func getAlphCode(c rune) int {
	if c-96 < 0 {
		return int(c) - 64 + 26
	} else {
		return int(c) - 96
	}
}

func part1(line string) int {
	first := line[0 : len(line)/2]
	second := line[len(line)/2:]
	// TODO: calc with set
	for _, c := range first {
		for _, c2 := range second {
			if c == c2 {
				return getAlphCode(c)
			}
		}
	}
	return 0
}

func part2(lines []string) int {
	for i := 0; i < len(lines[0]); i++ {
		for j := 0; j < len(lines[1]); j++ {
			for k := 0; k < len(lines[2]); k++ {
				if lines[0][i] == lines[1][j] && lines[1][j] == lines[2][k] {
					return getAlphCode(rune(lines[0][i]))
				}
			}
		}
	}
	return 0
}

func main() {
	start := time.Now()

	input, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(input), "\n")

	sum := 0
	for _, line := range lines {
		sum += part1(line)
	}

	for i := 0; i < len(lines); i++ {
		threeLines := lines[i : i+3]
		i += 2
		sum += part2(threeLines)
	}

	fmt.Println(sum)
	end := time.Now()
	fmt.Println(end.Sub(start))
}
