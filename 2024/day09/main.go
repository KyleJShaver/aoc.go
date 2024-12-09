package main

import (
	_ "embed"
	"fmt"
	"slices"

	common "github.com/KyleJShaver/aoc.go/common"
)

var year, day int = 2024, 9

//go:embed input.txt
var input string
var logger = common.Logger()

type MemoryBlock struct {
	Id     int
	Length int
}

func extractBlocks(input string, grouped bool) (blocks []*MemoryBlock) {
	blocks = make([]*MemoryBlock, 0)
	emptyBlock := MemoryBlock{Id: -1, Length: 1}
	maxId := 0
	for i := 0; i < len(input); i += 2 {
		block := MemoryBlock{Id: maxId, Length: 1}
		blockLength := common.Atoi(input[i])
		if !grouped {
			for range blockLength {
				blocks = append(blocks, &block)
			}
		} else {
			blocks = append(blocks, &MemoryBlock{Id: maxId, Length: blockLength})
		}
		maxId += 1
		if i+1 == len(input) {
			continue
		}
		if emptyLength := common.Atoi(input[i+1]); emptyLength > 0 {
			if !grouped {
				for range emptyLength {
					blocks = append(blocks, &emptyBlock)
				}
			} else {
				blocks = append(blocks, &MemoryBlock{Id: -1, Length: emptyLength})
			}
		}
	}
	return blocks
}

func defrag(blocks []*MemoryBlock) {
	blocksLen := len(blocks)
	fullBlocks := len(common.FilterSlice(blocks, func(block *MemoryBlock) bool {
		return block.Id != -1
	}))
	for i := range fullBlocks {
		if blocks[i].Id < 0 {
			for rev := range blocksLen {
				j := blocksLen - 1 - rev
				if blocks[j].Id > 0 {
					blocks[i], blocks[j] = blocks[j], blocks[i]
					break
				}
			}
		}
	}
}

func condsolidateEmptyMemory(blocks []*MemoryBlock) {
	for i := len(blocks) - 2; i >= 1; i-- {
		if blocks[i].Id == -1 && blocks[i-1].Id == -1 {
			blocks[i-1].Length += blocks[i].Length
			blocks = append(blocks[:i], blocks[i+1:]...)
		}
	}
}

func trim(blocks []*MemoryBlock) []*MemoryBlock {
	i := len(blocks) - 1
	if blocks[i] == blocks[i-1] {
		return trim(blocks[:i])
	}
	return blocks
}

func defragFiles(blocks []*MemoryBlock) {
	slices.Reverse(blocks)
	for i := range blocks {
		if blocks[i].Id > 0 {
			for rev := range blocks {
				j := len(blocks) - 1 - rev
				if j == i {
					break
				}
				if blocks[j].Id < 0 && blocks[j].Length >= blocks[i].Length {
					if blocks[j].Length == blocks[i].Length {
						blocks[i], blocks[j] = blocks[j], blocks[i]
						condsolidateEmptyMemory(blocks)
						break
					}
					splitEmpty := []*MemoryBlock{{Id: -1, Length: blocks[i].Length}, {Id: -1, Length: blocks[j].Length - blocks[i].Length}}
					blocks[i], blocks[j] = splitEmpty[0], blocks[i]
					blocks = append(blocks[:j], append(splitEmpty[1:], blocks[j:]...)...)
					condsolidateEmptyMemory(blocks)
					break
				}
			}
		}
	}
	blocks = trim(blocks)
	slices.Reverse(blocks)
}

func checksum(blocks []*MemoryBlock) (c int) {
	c = 0
	pos := 0
	for _, block := range blocks {
		if block.Id < 0 {
			pos += block.Length
			continue
		}
		if block.Length == 1 {
			c += pos * block.Id
			pos += 1
		} else {
			for range block.Length {
				c += pos * block.Id
				pos += 1
			}
		}

	}
	return
}

func Sprint(blocks []*MemoryBlock) (s string) {
	s = ""
	for _, block := range blocks {
		strRep := fmt.Sprintf("%d", block.Id)
		if block.Id < 0 {
			strRep = "."
		}
		for range block.Length {
			s += strRep
		}
	}
	return
}

func Part1(input string) (result int) {
	blocks := extractBlocks(input, false)
	defrag(blocks)
	result = checksum(blocks)
	logger.Info(fmt.Sprintf("Part 1: %v", result))
	return
}

func Part2(input string) (result int) {
	blocks := extractBlocks(input, true)
	defragFiles(blocks)
	result = checksum(blocks)
	logger.Info(fmt.Sprintf("Part 2: %v", result))
	return
}

func Day() (part1, part2 int) {
	defer common.Timer(fmt.Sprintf("%d-12-%02d", year, day), logger.Info, false)()
	part1 = Part1(input)
	part2 = Part2(input)
	return
}

func main() {
	_, _ = Day()
}
