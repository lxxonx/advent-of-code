package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func createScreen() [][]bool {
	width := 50
	height := 6
	grid := make([][]bool, height)
	for i := range grid {
		grid[i] = make([]bool, width)
	}
	return grid
}

func litScreen(grid [][]bool, width, height int) [][]bool {
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			grid[i][j] = true
		}
	}
	return grid
}

func rotateRow(row []bool, amount int) []bool {
	rotated := make([]bool, len(row))
	for i := range row {
		rotated[(i+amount)%len(row)] = row[i]
	}
	return rotated
}

func rotateColumn(grid [][]bool, column, amount int) [][]bool {
	rotated := make([]bool, len(grid))
	for i := range grid {
		rotated[(i+amount)%len(grid)] = grid[i][column]
	}
	for i := range grid {
		grid[i][column] = rotated[i]
	}
	return grid
}

func rotate(grid [][]bool, direction string, index, amount int) [][]bool {
	if direction == "row" {
		grid[index] = rotateRow(grid[index], amount)
	} else {
		grid = rotateColumn(grid, index, amount)
	}
	return grid
}

func main() {
	screen := createScreen()

	if input, err := ioutil.ReadFile("input.txt"); err == nil {
		lines := strings.Split(string(input), "\n")
		for _, line := range lines {
			chars := strings.Split(line, " ")

			switch chars[0] {
			case "rect":
				dimensions := strings.Split(chars[1], "x")

				width := 0
				fmt.Sscanf(dimensions[0], "%d", &width)

				height := 0
				fmt.Sscanf(dimensions[1], "%d", &height)

				screen = litScreen(screen, width, height)

			case "rotate":
				direction := chars[1]

				index := 0
				fmt.Sscanf(strings.Split(chars[2], "=")[1], "%d", &index)

				amount := 0
				fmt.Sscanf(chars[4], "%d", &amount)

				screen = rotate(screen, direction, index, amount)
			}
		}

		litPixels := 0
		for _, row := range screen {
			for _, pixel := range row {
				if pixel {
					litPixels++
				}
			}
		}

		fmt.Println(litPixels)

		for _, row := range screen {
			for _, pixel := range row {
				if pixel {
					fmt.Print("#")
				} else {
					fmt.Print(" ")
				}
			}
			fmt.Println()
		}

	}
}
