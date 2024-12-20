package main

import (
	"testing"

	"github.com/KyleJShaver/aoc.go/common"
)

var testInput string = `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`

func TestPart1(t *testing.T) {
	tests := []common.AOCTest{
		{
			Expected: 14,
			Fn:       Part1,
			Label:    "Part1",
			Input:    testInput,
		},
		{
			Expected: 4,
			Fn:       Part1,
			Label:    "Part1",
			Input: `..........
..........
..........
....a.....
........a.
.....a....
..........
..........
..........
..........`,
		},
		{
			Expected: 4,
			Fn:       Part1,
			Label:    "Part1",
			Input: `..........
..........
..........
....a.....
........a.
.....a....
..........
......A...
..........
..........`,
		},
	}
	for _, tst := range tests {
		tst.Run(t)
	}
}

func TestPart2(t *testing.T) {
	tests := []common.AOCTest{
		{
			Expected: 0,
			Fn:       Part2,
			Label:    "Part2",
			Input:    testInput,
		},
	}
	for _, tst := range tests {
		tst.Run(t)
	}
}
