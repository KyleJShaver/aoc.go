package main

import (
	"testing"

	"github.com/KyleJShaver/aoc.go/common"
)

var testInput string = `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

func TestPart1(t *testing.T) {
	tests := []common.AOCTest{
		{
			Expected: 3749,
			Fn:       Part1,
			Label:    "Part1",
			Input:    testInput,
		},
		{
			Expected: 3749,
			Fn:       Part1,
			Label:    "Part1",
			Input:    testInput,
		},
	}
	for _, tst := range tests {
		tst.Run(t)
	}
}

func TestPart2(t *testing.T) {
	tests := []common.AOCTest{
		{
			Expected: 11387,
			Fn:       Part2,
			Label:    "Part2",
			Input:    testInput,
		},
	}
	for _, tst := range tests {
		tst.Run(t)
	}
}
