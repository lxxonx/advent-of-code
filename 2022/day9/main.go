package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type Coordination struct {
	x int
	y int
}

func (c *Coordination) move(direction string) {
	switch direction {
	case "U":
		c.y++
	case "D":
		c.y--
	case "R":
		c.x++
	case "L":
		c.x--
	}
}

func (tail *Coordination) IsAdjacent(head Coordination) bool {
	if tail.x == head.x && tail.y == head.y {
		return true
	}
	if tail.x == head.x && (tail.y == head.y+1 || tail.y == head.y-1) {
		return true
	}
	if tail.y == head.y && (tail.x == head.x+1 || tail.x == head.x-1) {
		return true
	}
	if (tail.y == head.y+1 || tail.y == head.y-1) && (tail.x == head.x+1 || tail.x == head.x-1) {
		return true
	}
	return false
}

func (tail *Coordination) FollowHead(head Coordination) Coordination {
	dis := math.Max(math.Abs(float64(tail.x-head.x)), math.Abs(float64(tail.y-head.y)))

	if dis > 1 {
		dirX := head.x - tail.x
		if math.Abs(float64(dirX)) == 2 {
			tail.x += dirX / 2
		} else {
			tail.x += dirX
		}

		dirY := head.y - tail.y
		if math.Abs(float64(dirY)) == 2 {
			tail.y += dirY / 2
		} else {
			tail.y += dirY
		}
	}

	return *tail
}

func IsContain(visited []Coordination, coord Coordination) bool {
	for _, v := range visited {
		if v.x == coord.x && v.y == coord.y {
			return true
		}
	}
	return false
}

func part1(lines []string) {
	head := Coordination{0, 0}
	tail := Coordination{0, 0}

	visited := []Coordination{tail}

	for _, line := range lines {
		instruction := strings.Split(line, " ")
		direction := instruction[0]
		distance, _ := strconv.Atoi(instruction[1])
		for i := 0; i < distance; i++ {
			head.move(direction)
			fmt.Print("head", head)
			if !tail.IsAdjacent(head) {
				visit := tail.FollowHead(head)
				fmt.Print("tail", visit)
				if !IsContain(visited, visit) {
					visited = append(visited, visit)
				}
			}
			fmt.Println()
		}
	}

	fmt.Println(visited)
	fmt.Println(len(visited))
}

func part2(lines []string) {
	knots := []Coordination{}

	for i := 0; i < 10; i++ {
		knots = append(knots, Coordination{0, 0})
	}

	visited := []Coordination{knots[len(knots)-1]}

	for _, line := range lines {
		instruction := strings.Split(line, " ")
		direction := instruction[0]
		distance, _ := strconv.Atoi(instruction[1])
		for i := 0; i < distance; i++ {
			knots[0].move(direction)
			for j := 1; j < len(knots); j++ {
				if !knots[j].IsAdjacent(knots[j-1]) {
					knots[j].FollowHead(knots[j-1])
				}
			}

			tail := knots[len(knots)-1]
			if !IsContain(visited, tail) {
				visited = append(visited, tail)
			}

		}
	}

	fmt.Println(visited)
	fmt.Println(len(visited))
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")

	lines := strings.Split(string(input), "\n")

	part1(lines)

	part2(lines)
}
