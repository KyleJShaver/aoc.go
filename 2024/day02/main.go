package main

import (
	_ "embed"
	"fmt"

	common "github.com/KyleJShaver/aoc.go/common"
)

//go:embed input.txt
var input string
var InputInts = common.CachedListInts(" ")
var logger = common.Logger()

func reportIsSafe(report []int, dampener int) bool {
	direction := 0
	for pos, level := range report {
		if pos == 0 {
			direction = report[pos+1] - level
			continue
		}
		gap := level - report[pos-1]
		if (direction > 0) != (gap > 0) || gap <= -4 || gap >= 4 || gap == 0 {
			if dampener == 0 {
				return false
			}
			for i := range len(report) {
				new_report := make([]int, 0, len(report)-1)
				new_report = append(new_report, report[:i]...)
				new_report = append(new_report, report[i+1:]...)
				if reportIsSafe(new_report, dampener-1) {
					return true
				}
			}
			return false
		}
	}
	return true
}

func Part1(input string) (safe int) {
	reports := InputInts(input)
	safe = 0
	for _, report := range reports {
		if !reportIsSafe(report, 0) {
			continue
		}
		safe += 1
	}
	logger.Info(fmt.Sprintf("Part 1: %v", safe))
	return
}

func Part2(input string) (safe int) {
	reports := InputInts(input)
	safe = 0
	for _, report := range reports {
		if !reportIsSafe(report, 1) {
			continue
		}
		safe += 1
	}
	logger.Info(fmt.Sprintf("Part 2: %v", safe))
	return
}

func Day() (part1, part2 int) {
	defer common.Timer("2024-12-02", logger.Info, false)()
	part1 = Part1(input)
	part2 = Part2(input)
	return
}

func main() {
	_, _ = Day()
}
