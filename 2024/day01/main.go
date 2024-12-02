package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	common "github.com/KyleJShaver/aoc.go/common"
)

//go:embed input.txt
var input string
var SortedLists = CachedSortedLists()
var logger = common.Logger()

func CachedSortedLists() func(input string) [][]int {
	cache := make(map[string][][]int)
	return func(input string) (lists [][]int) {
		if retVal, ok := cache[input]; ok {
			return retVal
		}
		lines := strings.Split(input, "\n")
		logger.Debug(fmt.Sprintf("`input` has %v lines", len(lines)))
		lists = [][]int{make([]int, 0, len(lines)), make([]int, 0, len(lines))}
		for _, line := range lines {
			for pos, component := range strings.Split(line, "   ") {
				intVal, err := strconv.Atoi(component)
				common.Check(err)
				lists[pos] = append(lists[pos], intVal)
			}
		}
		for pos := range len(lists) {
			common.SortSlice(lists[pos])
		}
		cache[input] = lists
		return
	}
}

func Part1(input string) int {
	lists := SortedLists(input)
	difference := 0
	for pos, val := range lists[0] {
		pos_diff := val - lists[1][pos]
		if pos_diff < 0 {
			pos_diff *= -1
		}
		difference += pos_diff
	}
	logger.Info(fmt.Sprintf("Part 1: %v", difference))
	return difference
}

func Part2(input string) (score int) {
	lists := SortedLists(input)
	score = 0
	counts := make(map[int]int)
	for _, val := range lists[1] {
		counts[val] += 1
	}
	for _, val := range lists[0] {
		count, ok := counts[val]
		if ok {
			score += val * count
		}
	}
	logger.Info(fmt.Sprintf("Part 2: %v", score))
	return
}

func Day() (part1, part2 int) {
	defer common.Timer("2024-12-01", logger.Info, false)()
	part1 = Part1(input)
	part2 = Part2(input)
	return
}

func main() {
	_, _ = Day()
}
