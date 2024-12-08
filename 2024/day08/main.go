package main

import (
	_ "embed"
	"fmt"

	common "github.com/KyleJShaver/aoc.go/common"
)

var year, day int = 2024, 8

//go:embed input.txt
var input string
var logger = common.Logger()

func antennaLocations(grid [][]string) map[string][][]int {
	retVal := make(map[string][][]int)
	for y, row := range grid {
		for x, val := range row {
			if val == "." {
				continue
			}
			if _, ok := retVal[val]; !ok {
				retVal[val] = make([][]int, 0)
			}
			retVal[val] = append(retVal[val], []int{x, y})
		}
	}
	return retVal
}

func placeAntinodes(input string, antennaLocations map[string][][]int) map[string][][]int {
	antinodes := make(map[string][][]int)
	for antenna, locations := range antennaLocations {
		grid := common.Grid(input)
		for y, row := range grid {
			for x, val := range row {
				if val != antenna {
					grid[y][x] = "."
				}
			}
		}
		for _, seenAntinodes := range antinodes {
			for _, seenAntinode := range seenAntinodes {
				grid[seenAntinode[1]][seenAntinode[0]] = "#"
			}
		}
		antinodes[antenna] = make([][]int, 0)
		for pos, loc1 := range locations {
			for _, loc2 := range locations[pos+1:] {
				dist := []int{loc1[0] - loc2[0], loc1[1] - loc2[1]}
				loc1Antinode := []int{loc1[0] + dist[0], loc1[1] + dist[1]}
				loc2Antinode := []int{loc2[0] - dist[0], loc2[1] - dist[1]}
				for _, antinode := range [][]int{loc1Antinode, loc2Antinode} {
					if antinode[0] >= len(grid) || antinode[0] < 0 {
						continue
					}
					if antinode[1] >= len(grid[0]) || antinode[1] < 0 {
						continue
					}
					if grid[antinode[1]][antinode[0]] == "." {
						grid[antinode[1]][antinode[0]] = "#"
						antinodes[antenna] = append(antinodes[antenna], []int{antinode[0], antinode[1]})
					}
				}
			}
		}
	}
	return antinodes
}

func Part1(input string) (result int) {
	result = 0
	grid := common.Grid(input)
	antennaLocations := antennaLocations(grid)
	antinodeMap := placeAntinodes(input, antennaLocations)
	for _, antinodes := range antinodeMap {
		result += len(antinodes)
	}
	logger.Info(fmt.Sprintf("Part 1: %v", result))
	return
}

func Part2(input string) (result int) {
	result = 0
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
