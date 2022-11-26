package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func findAbba(str string) bool {
	for i := 0; i < len(str)-3; i++ {
		if str[i] == str[i+1] {
			continue
		}
		if str[i+1] != str[i+2] {
			continue
		}
		if str[i] != str[i+3] {
			continue
		}
		return true
	}
	return false
}

func findAba(str string) []string {
	aba := make([]string, 0)
	for i := 0; i < len(str)-2; i++ {
		if str[i] == str[i+1] {
			continue
		}
		if str[i] != str[i+2] {
			continue
		}
		aba = append(aba, str[i:i+3])
	}

	return aba
}

func findBab(str string, bab string) bool {
	return strings.Contains(str, bab)
}

func separateHypernet(str string) []string {
	replaced := strings.ReplaceAll(str, "[", " ")
	replaced = strings.ReplaceAll(replaced, "]", " ")

	return strings.Split(replaced, " ")
}

func part1(str string) bool {
	strArr := separateHypernet(str)
	res := false
	for i, s := range strArr {
		if i%2 == 1 {
			if findAbba(s) {
				return false
			}
		} else {
			if findAbba(s) {
				res = true
			}
		}
	}
	return res
}

func part2(str string) bool {
	strArr := separateHypernet(str)
	abas := make([]string, 0)

	for i, s := range strArr {
		if i%2 == 0 {
			b := findAba(s)
			abas = append(abas, b...)
		}
	}
	for i, s := range strArr {
		if i%2 == 1 {
			for _, b := range abas {
				if findBab(s, b[1:2]+b[0:1]+b[1:2]) {
					return true
				}
			}
		}
	}
	return false
}

func main() {
	count := 0
	if input, err := ioutil.ReadFile("input.txt"); err == nil {
		lines := strings.Split(string(input), "\n")
		for _, line := range lines {
			ok := part2(line)
			fmt.Printf(" %v\n", ok)
			if ok {
				count++
			}
		}
	}

	fmt.Println(count)
}
