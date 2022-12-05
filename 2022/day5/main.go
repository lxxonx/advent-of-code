package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

type Stack []string

func (s *Stack) Push(v string) {
	*s = append(*s, v)
}

func (s *Stack) PushArr(v []string) {
	*s = append(*s, v...)
}

func (s *Stack) Pop() string {
	if len(*s) == 0 {
		return ""
	}
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}

func (s *Stack) Peek() string {
	if len(*s) == 0 {
		return ""
	}
	return (*s)[len(*s)-1]
}

func (s *Stack) Slice(start, end int) Stack {
	return (*s)[start:end]
}

func part1(lines []string, arr []Stack) {

	for _, line := range lines {
		chars := strings.Split(line, " ")
		move, _ := strconv.Atoi(chars[1])
		from, _ := strconv.Atoi(chars[3])
		to, _ := strconv.Atoi(chars[5])

		for i := 0; i < move; i++ {
			arr[to-1].Push(arr[from-1].Pop())
		}
	}

	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i].Peek())
	}
	fmt.Println()
}

func part2(lines []string, arr []Stack) {
	for _, line := range lines {
		chars := strings.Split(line, " ")
		move, _ := strconv.Atoi(chars[1])
		from, _ := strconv.Atoi(chars[3])
		to, _ := strconv.Atoi(chars[5])

		a := arr[from-1].Slice(len(arr[from-1])-move, len(arr[from-1]))
		arr[from-1] = arr[from-1].Slice(0, len(arr[from-1])-move)

		arr[to-1].PushArr(a)
	}

	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i].Peek())
	}
	fmt.Println()
}

func main() {
	start := time.Now()
	input, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(input), "\n")

	s1 := Stack(strings.Split("BQC", ""))
	s2 := Stack(strings.Split("RQWZ", ""))
	s3 := Stack(strings.Split("BMRLV", ""))
	s4 := Stack(strings.Split("CZHVTW", ""))
	s5 := Stack(strings.Split("DZHBNVG", ""))
	s6 := Stack(strings.Split("HNPCJFVQ", ""))
	s7 := Stack(strings.Split("DGTRWZS", ""))
	s8 := Stack(strings.Split("CGMNBWZP", ""))
	s9 := Stack(strings.Split("NJBMWQFP", ""))

	arr := []Stack{s1, s2, s3, s4, s5, s6, s7, s8, s9}

	// Part 1
	// part1(lines, arr)

	// Part 2
	part2(lines, arr)

	fmt.Println(time.Since(start))
}
