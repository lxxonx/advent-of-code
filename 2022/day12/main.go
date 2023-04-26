package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

type coord struct {
	x, y int
}

func sol1(grid [][]rune, start, end coord) {
	visited := make(map[coord]bool)
	toVisit := []coord{start}
	distanceFromStart := map[coord]int{start: 0}

	for {
		currentPoint := toVisit[0]
		visited[currentPoint] = true
		toVisit = toVisit[1:]

		if currentPoint == end {
			fmt.Println(distanceFromStart[end])
			break
		}

		for _, near := range [][]int{{1, 0}, {0, -1}, {-1, 0}, {0, 1}} {
			j, i := near[1], near[0]
			nextPoint := coord{currentPoint.x + j, currentPoint.y + i}
			if !visited[nextPoint] && nextPoint.x >= 0 && nextPoint.y >= 0 &&
				nextPoint.x < len(grid[0]) && nextPoint.y < len(grid) &&
				(grid[nextPoint.y][nextPoint.x]-grid[currentPoint.y][currentPoint.x] <= 1) {
				if distanceFromStart[nextPoint] == 0 {
					toVisit = append(toVisit, nextPoint)
					distanceFromStart[nextPoint] = distanceFromStart[currentPoint] + 1
				}
				if distanceFromStart[nextPoint] >= distanceFromStart[currentPoint]+1 {
					distanceFromStart[nextPoint] = distanceFromStart[currentPoint] + 1
				}
			}
		}
		sort.Slice(toVisit, func(i, j int) bool {
			return distanceFromStart[toVisit[i]] < distanceFromStart[toVisit[j]]
		})
	}
}

func sol2(grid [][]rune, end coord) {
	starts := make([]coord, 0)
	candidates := make([]int, 0)

	for y, line := range grid {
		for x, elevation := range line {
			if elevation == 'a' {
				starts = append(starts, coord{x, y})
			}
		}
	}

	for _, start := range starts {
		visited := make(map[coord]bool)
		toVisit := []coord{start}
		distanceFromStart := map[coord]int{start: 0}
		for {
			if len(toVisit) == 0 {
				break
			}
			currentPoint := toVisit[0]
			visited[currentPoint] = true
			toVisit = toVisit[1:]

			if currentPoint == end {
				candidates = append(candidates, distanceFromStart[end])
				break
			}

			for _, near := range [][]int{{1, 0}, {0, -1}, {-1, 0}, {0, 1}} {
				j, i := near[1], near[0]
				nextPoint := coord{currentPoint.x + j, currentPoint.y + i}
				if !visited[nextPoint] && nextPoint.x >= 0 && nextPoint.y >= 0 &&
					nextPoint.x < len(grid[0]) && nextPoint.y < len(grid) &&
					(grid[nextPoint.y][nextPoint.x]-grid[currentPoint.y][currentPoint.x] <= 1) {
					if distanceFromStart[nextPoint] == 0 {
						toVisit = append(toVisit, nextPoint)
						distanceFromStart[nextPoint] = distanceFromStart[currentPoint] + 1
					}
					if distanceFromStart[nextPoint] >= distanceFromStart[currentPoint]+1 {
						distanceFromStart[nextPoint] = distanceFromStart[currentPoint] + 1
					}
				}
			}
			sort.Slice(toVisit, func(i, j int) bool {
				return distanceFromStart[toVisit[i]] < distanceFromStart[toVisit[j]]
			})
		}
	}

	min := math.MaxInt64
	for _, candidate := range candidates {
		if candidate < min {
			min = candidate
		}
	}
	fmt.Println(min)
}

func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	grid := make([][]rune, 0)
	var start, end coord

	for sc.Scan() {
		var line []rune
		for i, elevation := range sc.Text() {
			if elevation == 'S' {
				start = coord{i, len(grid)}
				elevation = 'a'
			}
			if elevation == 'E' {
				end = coord{i, len(grid)}
				elevation = 'z'
			}
			line = append(line, elevation)
		}
		grid = append(grid, line)
	}
	sol1(grid, start, end)
	sol2(grid, end)
}
