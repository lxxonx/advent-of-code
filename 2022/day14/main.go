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

func DrawWall(coords []Coordination, walls map[Coordination]rune) map[Coordination]rune {
	for {
		if len(coords) == 1 {
			break
		}
		curr := coords[0]
		next := coords[1]
		coords = coords[1:]

		for curr.x != next.x || curr.y != next.y {
			walls[curr] = '#'
			if curr.x < next.x {
				curr.x++
			}
			if curr.x > next.x {
				curr.x--
			}
			if curr.y < next.y {
				curr.y++
			}
			if curr.y > next.y {
				curr.y--
			}
		}
		walls[curr] = '#'
	}
	return walls
}

func PourSand(walls map[Coordination]rune) int {
	isFalling := false
	unit := 0

	bottom := 0

	for wall := range walls {
		if wall.y > bottom {
			bottom = wall.y
		}
	}

	for !isFalling {
		sand := Coordination{500, 0}
		for {
			walls[sand] = '"'
			if sand.y+1 > bottom {
				isFalling = true
				break
			}
			if walls[Coordination{sand.x, sand.y + 1}] < '#' {
				sand.y++
			} else if walls[Coordination{sand.x - 1, sand.y + 1}] < '#' {
				sand.y++
				sand.x--
			} else if walls[Coordination{sand.x + 1, sand.y + 1}] < '#' {
				sand.y++
				sand.x++
			} else {
				walls[sand] = 'o'
				unit++
				break
			}
		}
	}

	return unit
}

func PourSand2(walls map[Coordination]rune) int {
	unit := 0

	bottom := 0
	minX := math.MaxInt
	maxX := math.MinInt

	for wall := range walls {
		if wall.y > bottom {
			bottom = wall.y
		}
		if wall.x < minX {
			minX = wall.x
		}
		if wall.x > maxX {
			maxX = wall.x
		}
	}
	bottom = bottom + 2
	for i := minX - 500; i < maxX+500; i++ {
		walls[Coordination{i, bottom}] = '#'
	}

	for {
		sand := Coordination{500, 0}
		if walls[sand] == 'o' {
			break
		}
		for {
			walls[sand] = '"'
			if walls[Coordination{sand.x, sand.y + 1}] < '#' {
				sand.y++
			} else if walls[Coordination{sand.x - 1, sand.y + 1}] < '#' {
				sand.y++
				sand.x--
			} else if walls[Coordination{sand.x + 1, sand.y + 1}] < '#' {
				sand.y++
				sand.x++
			} else {
				walls[sand] = 'o'
				unit++
				break
			}
		}
	}
	return unit
}

func main() {

	walls := make(map[Coordination]rune)

	sample, _ := ioutil.ReadFile("input.txt")

	lines := strings.Split(string(sample), "\n")

	for _, line := range lines {
		coords := strings.Split(line, " -> ")
		Coordinations := []Coordination{}

		for _, coord := range coords {
			x, _ := strconv.Atoi(strings.Split(coord, ",")[0])
			y, _ := strconv.Atoi(strings.Split(coord, ",")[1])
			Coordinations = append(Coordinations, Coordination{x, y})
		}
		walls = DrawWall(Coordinations, walls)
	}

	sol1 := PourSand(walls)

	fmt.Println("Part 1:", sol1)

	sol2 := PourSand2(walls)

	fmt.Println("Part 2:", sol1+sol2)
}
