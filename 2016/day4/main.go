package main

import (
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type runeCount struct {
	char  string
	count int
}

func part1() {
	sum := 0
	if input, err := ioutil.ReadFile("input.txt"); err == nil {
		lines := strings.Split(string(input), "\n")
		for _, line := range lines {
			c := make([]runeCount, 0)
			chunk := strings.Split(string(line), "-")
			id := strings.Split(chunk[len(chunk)-1], "[")[0]
			checksum := strings.Split(chunk[len(chunk)-1], "[")[1][:5]

			chars := strings.Join(chunk[:len(chunk)-1], "")
			for _, char := range chars {
				found := false
				for i, r := range c {
					if r.char == string(char) {
						c[i].count++
						found = true
					}
				}
				if !found {
					c = append(c, runeCount{string(char), 1})
				}
			}
			sort.SliceStable(c, func(i, j int) bool {
				if c[i].count == c[j].count {
					return c[i].char < c[j].char
				}
				return c[i].count > c[j].count
			})
			isValid := false
			for i, r := range c {
				if i >= 5 {
					isValid = true
					break
				}
				if string(r.char) != string(checksum[i]) {
					break
				}
			}
			if isValid {
				num, err := strconv.Atoi(id)
				if err != nil {
					panic(err)
				}
				sum += num
			}

		}
	}

	println(sum)
}

func part2() {
	if input, err := ioutil.ReadFile("input.txt"); err == nil {
		lines := strings.Split(string(input), "\n")
		for _, line := range lines {
			chunk := strings.Split(string(line), "-")
			id := strings.Split(chunk[len(chunk)-1], "[")[0]
			chars := strings.Join(chunk[:len(chunk)-1], " ")

			shift, err := strconv.Atoi(id)
			if err != nil {
				panic(err)
			}
			shift = shift % 26
			shifted := ""
			for _, char := range chars {
				if char == ' ' {
					shifted += " "
					continue
				}
				if int(char)+shift > 122 {
					shifted += string(rune(int(char) + shift - 26))
				} else {
					shifted += string(rune(int(char) + shift))
				}
			}

			if strings.Contains(shifted, "north") {
				println(shifted, id)
			}

		}
	}
}

func main() {
	// part1()
	part2()
}
