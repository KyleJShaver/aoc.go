package main

import (
	"testing"

	"github.com/KyleJShaver/aoc.go/common"
)

func TestPart1(t *testing.T) {
	tests := []common.AOCTest{
		{
			Expected: 161,
			Fn:       Part1,
			Label:    "Part1",
			Input:    "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))",
		},
	}
	for _, tst := range tests {
		tst.Run(t)
	}
}

func TestPart2(t *testing.T) {
	tests := []common.AOCTest{
		{
			Expected: 48,
			Fn:       Part2,
			Label:    "Part2",
			Input:    "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
		},
		{
			Expected: 136,
			Fn:       Part2,
			Label:    "Part2",
			Input:    "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](do()mul(11,8)undo()?mul(8,5))",
		},
		{
			Expected: 96,
			Fn:       Part2,
			Label:    "Part2",
			Input:    "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](do()mul(11,8)undo()don't()?mul(8,5))",
		},
	}
	for _, tst := range tests {
		tst.Run(t)
	}
}
