package main

import (
	_ "embed"
	"fmt"
	"strings"

	common "github.com/KyleJShaver/aoc.go/common"
)

var year, day int = 2024, 06

//go:embed input.txt
var input string
var logger = common.Logger()

func guardPos(grid [][]string) []int {
	for y, row := range grid {
		for x, val := range row {
			if strings.Contains("^>v<", val) {
				return []int{x, y}
			}
		}
	}
	return nil
}

func guardDirection(grid [][]string, pos []int) (direction []int) {
	direction = nil
	switch grid[pos[1]][pos[0]] {
	case "^":
		direction = []int{0, -1}
	case ">":
		direction = []int{1, 0}
	case "v":
		direction = []int{0, 1}
	case "<":
		direction = []int{-1, 0}
	}
	return
}

func rotate(direction []int) []int {
	newDirection := []int{direction[1], direction[0]}
	if newDirection[0] != 0 {
		newDirection[0] *= -1
	}
	return newDirection
}

const (
	outOfBounds = iota
	obstruction = iota
	clear       = iota
	cycle       = iota
)

func nextPos(grid [][]string, pos, direction []int) (nextPos []int, content int) {
	nextPos = []int{pos[0] + direction[0], pos[1] + direction[1]}
	content = clear
	if nextPos[0] < 0 || nextPos[0] >= len(grid[0]) || nextPos[1] < 0 || nextPos[1] >= len(grid) {
		content = outOfBounds
		return
	}
	if grid[nextPos[1]][nextPos[0]] == "#" {
		content = obstruction
	}
	if grid[nextPos[1]][nextPos[0]] == "B" {
		content = cycle
	}
	return
}

var gridMarkers = map[string]string{
	".": "X",
	"X": "C",
	"C": "V",
	"V": "B",
}

func Part1(input string) (result int) {
	result = 0
	grid := common.Grid(input)
	pos := guardPos(grid)
	direction := guardDirection(grid, pos)
	np, c := nextPos(grid, pos, direction)
	for c != outOfBounds {
		if c == clear {
			grid[np[1]][np[0]] = grid[pos[1]][pos[0]]
			grid[pos[1]][pos[0]] = "X"
			pos = np
			np, c = nextPos(grid, pos, direction)
		}
		if c == obstruction {
			direction = rotate(direction)
			np, c = nextPos(grid, pos, direction)
		}
		if c == cycle {
			common.Check(fmt.Errorf("cycle detected in part 1"))
		}
	}
	grid[pos[1]][pos[0]] = "X"
	for _, row := range grid {
		for _, val := range row {
			if val == "X" {
				result += 1
			}
		}
	}
	logger.Info(fmt.Sprintf("Part 1: %v", result))
	return
}

func part2Obstacle(input string, x, y int) int {
	grid := common.Grid(input)
	grid[y][x] = "#"
	pos := guardPos(grid)
	direction := guardDirection(grid, pos)
	grid[pos[1]][pos[0]] = "."
	np, c := nextPos(grid, pos, direction)
	for c != outOfBounds {
		if c == clear {
			grid[pos[1]][pos[0]] = gridMarkers[grid[pos[1]][pos[0]]]
			pos = np
			np, c = nextPos(grid, pos, direction)
		}
		if c == obstruction {
			direction = rotate(direction)
			np, c = nextPos(grid, pos, direction)
		}
		if c == cycle {
			return 1
		}
	}
	return 0
}

// Brute force is still force.
func Part2(input string) (result int) {
	result = 0
	grid := common.Grid(input)
	for y, row := range grid {
		for x, val := range row {
			if strings.Contains("#^", val) {
				continue
			}
			result += part2Obstacle(input, x, y)
		}
	}
	logger.Info(fmt.Sprintf("Part 2: %v", result))
	return
}

func Day() (part1, part2 int) {
	defer common.Timer(fmt.Sprintf("%d-12-%02d", year, day), logger.Info, false)()
	part1 = Part1(input)
	part2 = Part2(input)
	return
}

func main() {
	_, _ = Day()
}
