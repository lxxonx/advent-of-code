package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type charCount struct {
	char  string
	count int
}

func part1(counter map[int][]charCount) {
	str := make([]string, len(counter))

	for i, c := range counter {
		sort.SliceStable(c, func(i, j int) bool {
			if c[i].count == c[j].count {
				return c[i].char < c[j].char
			}
			return c[i].count > c[j].count
		})

		str[i] = c[0].char
	}

	fmt.Println(strings.Join(str, ""))
}

func part2(counter map[int][]charCount) {
	str := make([]string, len(counter))

	for i, c := range counter {
		sort.SliceStable(c, func(i, j int) bool {
			if c[i].count == c[j].count {
				return c[i].char > c[j].char
			}
			return c[i].count < c[j].count
		})

		str[i] = c[0].char
	}

	fmt.Println(strings.Join(str, ""))
}

func main() {
	counter := make(map[int][]charCount)
	if input, err := ioutil.ReadFile("input.txt"); err == nil {
		lines := strings.Split(string(input), "\n")
		for _, line := range lines {
			chars := strings.Split(line, "")
			for i, char := range chars {
				if _, ok := counter[i]; !ok {
					counter[i] = make([]charCount, 0)
				}
				found := false
				for j, c := range counter[i] {
					if c.char == char {
						counter[i][j].count++
						found = true
						break
					}
				}
				if !found {
					counter[i] = append(counter[i], charCount{char, 1})
				}
			}
		}
	}
	// part1(counter)
	part2(counter)
}
