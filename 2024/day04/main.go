package main

import (
	_ "embed"
	"fmt"

	common "github.com/KyleJShaver/aoc.go/common"
)

//go:embed input.txt
var input string
var logger = common.Logger()
var directionMap = map[string][]int{
	"Up":    {0, -1},
	"Down":  {0, 1},
	"Left":  {-1, 0},
	"Right": {1, 0},
}

func init() {
	for _, vertical := range []string{"Up", "Down"} {
		for _, horizontal := range []string{"Left", "Right"} {
			directionMap[vertical+horizontal] = []int{directionMap[horizontal][0], directionMap[vertical][1]}
		}
	}
}

func checkDirection(grid [][]string, search string, direction []int, x, y int) bool {
	if len(grid) == 0 {
		return false
	}
	if y >= len(grid) || x >= len(grid[0]) || y < 0 || x < 0 {
		return false
	}
	if (y < len(search)-1 && direction[1] < 0) || (y >= len(grid) && direction[1] > 0) {
		return false
	}
	if (x < len(search)-1 && direction[0] < 0) || (x >= len(grid[0]) && direction[0] > 0) {
		return false
	}
	if grid[y][x] != search[0:1] {
		return false
	} else if len(search) == 1 {
		return true
	}
	return checkDirection(grid, search[1:], direction, x+direction[0], y+direction[1])
}

func checkDirections(grid [][]string, search string) (directions [][][]string) {
	for y, row := range grid {
		directions = append(directions, make([][]string, 0))
		for x := range row {
			directions[y] = append(directions[y], make([]string, 0))
			for directionKey, direction := range directionMap {
				if !checkDirection(grid, search, direction, x, y) {
					continue
				}
				directions[y][x] = append(directions[y][x], directionKey)
			}
		}
	}
	return
}

func checkDirectionsX(grid [][]string, search string) (directions [][][]string) {
	for y, row := range grid {
		directions = append(directions, make([][]string, 0))
		for x := range row {
			directions[y] = append(directions[y], make([]string, 0))
			directionKey := "DownRight"
			direction := directionMap[directionKey]
			searchesMuts := []string{search, common.ReverseString(search)}
			found := false
			for _, searchMut := range searchesMuts {
				found = found || checkDirection(grid, searchMut, direction, x, y)
			}
			if !found {
				continue
			}
			direction = directionMap["DownLeft"]
			found = false
			for _, searchMut := range searchesMuts {
				found = found || checkDirection(grid, searchMut, direction, x+2, y)
			}
			if !found {
				continue
			}
			directions[y][x] = append(directions[y][x], directionKey)
		}
	}
	return
}

func Part1(input string) (result int) {
	result = 0

	grid := common.GridCached(input)
	for _, row := range checkDirections(grid, "XMAS") {
		for _, val := range row {
			result += len(val)
		}
	}

	logger.Info(fmt.Sprintf("Part 1: %v", result))
	return
}

func Part2(input string) (result int) {
	result = 0
	grid := common.GridCached(input)

	for _, row := range checkDirectionsX(grid, "MAS") {
		for _, val := range row {
			result += len(val)
		}
	}
	logger.Info(fmt.Sprintf("Part 2: %v", result))
	return
}

func Day() (part1, part2 int) {
	defer common.Timer("2024-12-04", logger.Info, false)()
	part1 = Part1(input)
	part2 = Part2(input)
	return
}

func main() {
	_, _ = Day()
}
