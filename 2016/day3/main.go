package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func part1(line string) bool {
	strs := strings.Fields(line)
	nums := make([]int, 0)

	for _, i := range strs {
		num, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		nums = append(nums, num)
	}
	sort.Ints(nums)
	if nums[0]+nums[1] > nums[2] {
		return true
	} else {
		return false
	}
}

func part2(nums []int) bool {
	numsCopy := make([]int, len(nums))
	copy(numsCopy, nums)
	sort.Ints(numsCopy)
	result := false
	if numsCopy[0]+numsCopy[1] > numsCopy[2] {
		result = true
	}

	fmt.Println(nums, result)

	return result
}

func main() {
	count := 0

	if input, err := ioutil.ReadFile("input.txt"); err == nil {
		lines := strings.Split(string(input), "\n")
		cols := [][]int{{}, {}, {}}
		for _, line := range lines {
			strs := strings.Fields(line)
			for i, j := range strs {
				num, err := strconv.Atoi(j)
				if err != nil {
					panic(err)
				}
				if len(cols[i]) == 0 {
					cols[i] = make([]int, 0)
				}
				cols[i] = append(cols[i], num)
				if len(cols[i]) == 3 {
					if part2(cols[i]) {
						count++
					}
					cols[i] = []int{}
				}
			}
		}
	}

	fmt.Println(count)
}
