package main

import (
	_ "embed"
	"fmt"

	common "github.com/KyleJShaver/aoc.go/common"
)

//go:embed input.txt
var input string
var logger = common.Logger()

func Part1(input string) (result int) {
	result = 0
	logger.Info(fmt.Sprintf("Part 1: %v", result))
	return
}

func Part2(input string) (result int) {
	result = 0
	logger.Info(fmt.Sprintf("Part 2: %v", result))
	return
}

func Day() (part1, part2 int) {
	defer common.Timer("2024-12-##", logger.Info, false)()
	part1 = Part1(input)
	part2 = Part2(input)
	return
}

func main() {
	_, _ = Day()
}