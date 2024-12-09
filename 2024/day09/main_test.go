package main

import (
	"testing"

	"github.com/KyleJShaver/aoc.go/common"
)

var testInput string = `2333133121414131402`

func TestPart1(t *testing.T) {
	tests := []common.AOCTest{
		{
			Expected: 1928,
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
			Expected: 2858,
			Fn:       Part2,
			Label:    "Part2",
			Input:    testInput,
		},
	}
	for _, tst := range tests {
		tst.Run(t)
	}
}
