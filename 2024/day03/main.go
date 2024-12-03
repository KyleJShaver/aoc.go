package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	common "github.com/KyleJShaver/aoc.go/common"
)

//go:embed input.txt
var input string
var logger = common.Logger()
var multRe = regexp.MustCompile(`mul\((\d+),(\d+)\)`)

func sumMulCommands(input string) int {
	result := 0
	for _, submatches := range multRe.FindAllStringSubmatch(input, -1) {
		left, err := strconv.Atoi(submatches[1])
		common.Check(err)
		right, err := strconv.Atoi(submatches[2])
		common.Check(err)
		result += right * left
	}
	return result
}

func Part1(input string) (result int) {
	result = sumMulCommands(input)
	logger.Info(fmt.Sprintf("Part 1: %v", result))
	return
}

// get the next `don't()` to `do()` range
func dontDo(input string) (dontPos, doPos int) {
	dontPos = strings.Index(input, "don't()")
	if dontPos < 0 {
		return -1, -1
	}
	doPos = strings.Index(input[dontPos:], "do()")
	if doPos < 0 {
		return dontPos, len(input)
	} else {
		doPos += dontPos
	}
	return
}

func removeDisabledMemory(input string) (enabledMemory string) {
	enabledMemory = input
	for {
		dontPos, doPos := dontDo(enabledMemory)
		if dontPos < 0 || doPos < 0 {
			break
		}
		enabledMemory = enabledMemory[:dontPos] + enabledMemory[doPos:]
	}
	return
}

func Part2(input string) (result int) {
	result = 0
	inputDisabledRemoved := removeDisabledMemory(input)
	result = sumMulCommands(inputDisabledRemoved)
	logger.Info(fmt.Sprintf("Part 2: %v", result))
	return
}

func Day() (part1, part2 int) {
	defer common.Timer("2024-12-03", logger.Info, false)()
	part1 = Part1(input)
	part2 = Part2(input)
	return
}

func main() {
	_, _ = Day()
}
