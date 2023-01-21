package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func splitVertical(square [][]int, x int) [][]int {
	vertical := []int{}
	for _, line := range square {
		vertical = append(vertical, line[x])
	}

	return [][]int{vertical, square[x]}
}

func hasBlockingTree(left []int, target int) bool {
	for _, num := range left {
		if num >= target {
			return false
		}
	}
	return true
}

func part1(square [][]int) {
	count := 0

	for i, line := range square {
		for j := range line {
			if i == 0 || j == 0 || i == len(square)-1 || j == len(line)-1 {
				count++
				continue
			}

			left := line[:j]
			right := line[j+1:]
			if hasBlockingTree(left, line[j]) {
				count++
				continue
			}
			if hasBlockingTree(right, line[j]) {
				count++
				continue
			}

			vertical := splitVertical(square, j)
			up := vertical[0][:i]
			down := vertical[0][i+1:]
			if hasBlockingTree(up, line[j]) {
				count++
				continue
			}
			if hasBlockingTree(down, line[j]) {
				count++
				continue
			}
		}

	}
	fmt.Println(count)
}

func reverse(arr []int) []int {
	s := make([]int, len(arr))
	copy(s, arr)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func countSeenTrees(line []int, target int) int {
	count := 0
	for _, num := range line {
		count++
		if num >= target {
			break
		}
	}
	return count
}

func part2(square [][]int) {
	max := 0

	for i, line := range square {
		for j := range line {
			if i == 0 || j == 0 || i == len(square)-1 || j == len(line)-1 {
				continue
			}

			left := line[:j]
			left = reverse(left)
			right := line[j+1:]

			leftCount := countSeenTrees(left, line[j])
			rightCount := countSeenTrees(right, line[j])

			vertical := splitVertical(square, j)[0]
			up := vertical[:i]
			up = reverse(up)
			down := vertical[i+1:]

			upCount := countSeenTrees(up, line[j])
			downCount := countSeenTrees(down, line[j])

			count := leftCount * rightCount * upCount * downCount

			if int(math.Max(float64(count), float64(max))) == count {
				max = count
			}
		}

	}

	fmt.Println(max)
}

func main() {
	square := [][]int{}
	input, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(input), "\n")

	for _, line := range lines {
		square = append(square, []int{})
		chars := strings.Split(line, "")
		for _, char := range chars {
			num, _ := strconv.Atoi(char)
			square[len(square)-1] = append(square[len(square)-1], num)
		}
	}

	// part1(square)
	part2(square)

}
