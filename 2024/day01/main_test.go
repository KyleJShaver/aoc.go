package main

import (
	"testing"

	"github.com/KyleJShaver/aoc.go/common"
)

var testInput string = `3   4
4   3
2   5
1   3
3   9
3   3`

func TestPart1(t *testing.T) {
	tests := []common.AOCTest{
		{
			Expected: 11,
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
			Expected: 31,
			Fn:       Part2,
			Label:    "Part2",
			Input:    testInput,
		},
	}
	for _, tst := range tests {
		tst.Run(t)
	}
}
