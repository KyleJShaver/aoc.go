package main

import (
	"testing"

	"github.com/KyleJShaver/aoc.go/common"
)

var testInput string = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

func TestPart1(t *testing.T) {
	tests := []common.AOCTest{
		{
			Expected: 41,
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
			Expected: 6,
			Fn:       Part2,
			Label:    "Part2",
			Input:    testInput,
		},
	}
	for _, tst := range tests {
		tst.Run(t)
	}
}
