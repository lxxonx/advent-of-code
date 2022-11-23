package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type coord struct {
	x int
	y int
}

func move(coordinate [2]int, char string, face string) ([2]int, string) {
	char = strings.TrimSpace(char)

	dir := char[:1]
	dist, err := strconv.Atoi(char[1:])
	if err != nil {
		panic(err)
	}

	switch face {
	case "N":
		if dir == "R" {
			face = "E"
			coordinate[0] += dist
		} else {
			face = "W"
			coordinate[0] -= dist
		}
	case "E":
		if dir == "R" {
			face = "S"
			coordinate[1] -= dist
		} else {
			face = "N"
			coordinate[1] += dist
		}
	case "S":
		if dir == "R" {
			face = "W"
			coordinate[0] -= dist
		} else {
			face = "E"
			coordinate[0] += dist
		}
	case "W":
		if dir == "R" {
			face = "N"
			coordinate[1] += dist
		} else {
			face = "S"
			coordinate[1] -= dist
		}
	}
	return coordinate, face
}

func part1() {
	coord := [2]int{0, 0}
	face := "N"

	// Read input
	if input, err := ioutil.ReadFile("input.txt"); err == nil {
		chars := strings.Split(string(input), ",")
		for _, char := range chars {
			coord, face = move(coord, char, face)
		}
	}

	result := math.Abs(float64(coord[1])) + math.Abs(float64(coord[0]))
	fmt.Println(result)
}

func isVisited(cood [2]int, visited []coord) bool {
	for _, v := range visited {
		if v.x == cood[0] && v.y == cood[1] {
			return true
		}
	}
	return false
}

func mark(coordinate [2]int, char string, face string, visited []coord) (bool, [2]int, string, []coord) {
	char = strings.TrimSpace(char)

	dir := char[:1]
	dist, err := strconv.Atoi(char[1:])
	if err != nil {
		panic(err)
	}

	switch face {
	case "N":
		if dir == "R" {
			face = "E"

			for i := 1; i <= dist; i++ {
				isVis := isVisited([2]int{coordinate[0] + i, coordinate[1]}, visited)
				if !isVis {
					visited = append(visited, coord{coordinate[0] + i, coordinate[1]})
				} else {
					return true, [2]int{coordinate[0] + i, coordinate[1]}, face, visited
				}
			}
			coordinate[0] += dist

		} else {
			face = "W"
			for i := 1; i <= dist; i++ {
				isVis := isVisited([2]int{coordinate[0] - i, coordinate[1]}, visited)
				if !isVis {
					visited = append(visited, coord{coordinate[0] - i, coordinate[1]})
				} else {
					return true, [2]int{coordinate[0] - i, coordinate[1]}, face, visited
				}
			}
			coordinate[0] -= dist
		}
	case "E":
		if dir == "R" {
			face = "S"
			for i := 1; i <= dist; i++ {
				isVis := isVisited([2]int{coordinate[0], coordinate[1] - i}, visited)
				if !isVis {
					visited = append(visited, coord{coordinate[0], coordinate[1] - i})
				} else {
					return true, [2]int{coordinate[0], coordinate[1] - i}, face, visited
				}
			}
			coordinate[1] -= dist
		} else {
			face = "N"
			for i := 1; i <= dist; i++ {
				isVis := isVisited([2]int{coordinate[0], coordinate[1] + i}, visited)
				if !isVis {
					visited = append(visited, coord{coordinate[0], coordinate[1] + i})
				} else {
					return true, [2]int{coordinate[0], coordinate[1] + i}, face, visited
				}
			}
			coordinate[1] += dist
		}
	case "S":
		if dir == "R" {
			face = "W"

			for i := 1; i <= dist; i++ {
				isVis := isVisited([2]int{coordinate[0] - i, coordinate[1]}, visited)
				if !isVis {
					visited = append(visited, coord{coordinate[0] - i, coordinate[1]})
				} else {
					return true, [2]int{coordinate[0] - i, coordinate[1]}, face, visited
				}
			}
			coordinate[0] -= dist
		} else {
			face = "E"

			for i := 1; i <= dist; i++ {
				isVis := isVisited([2]int{coordinate[0] + i, coordinate[1]}, visited)
				if !isVis {
					visited = append(visited, coord{coordinate[0] + i, coordinate[1]})
				} else {
					return true, [2]int{coordinate[0] + i, coordinate[1]}, face, visited
				}
			}
			coordinate[0] += dist
		}
	case "W":
		if dir == "R" {
			face = "N"

			for i := 1; i <= dist; i++ {
				isVis := isVisited([2]int{coordinate[0], coordinate[1] + i}, visited)
				if !isVis {
					visited = append(visited, coord{coordinate[0], coordinate[1] + i})
				} else {
					return true, [2]int{coordinate[0], coordinate[1] + i}, face, visited
				}
			}
			coordinate[1] += dist
		} else {
			face = "S"
			for i := 1; i <= dist; i++ {
				isVis := isVisited([2]int{coordinate[0], coordinate[1] - i}, visited)
				if !isVis {
					visited = append(visited, coord{coordinate[0], coordinate[1] - i})
				} else {
					return true, [2]int{coordinate[0], coordinate[1] - i}, face, visited
				}
			}
			coordinate[1] -= dist
		}
	}

	return false, coordinate, face, visited
}

func part2() {
	cood := [2]int{0, 0}
	face := "N"
	visited := []coord{}
	found := false

	// Read input
	if input, err := ioutil.ReadFile("input.txt"); err == nil {
		chars := strings.Split(string(input), ",")
		for _, char := range chars {
			found, cood, face, visited = mark(cood, char, face, visited)
			if found {
				break
			}
		}
	}
	fmt.Println(cood)

	result := math.Abs(float64(cood[1])) + math.Abs(float64(cood[0]))
	fmt.Println(result)
}

func main() {
	part2()
}
