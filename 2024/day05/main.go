package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"

	common "github.com/KyleJShaver/aoc.go/common"
)

//go:embed input.txt
var input string
var logger = common.Logger()

type Node struct {
	value  int
	before []*Node
	after  []*Node
}

func newNode(value int) (n *Node) {
	return &Node{value: value, before: make([]*Node, 0), after: make([]*Node, 0)}
}

func inputParser(input string) (updates [][]*Node, e error) {
	sections := strings.Split(input, "\n\n")
	orderingRules, pageUpdates := sections[0], sections[1]
	rules, updates, e := make(map[int]*Node), make([][]*Node, 0), nil
	for _, orderingRule := range strings.Split(orderingRules, "\n") {
		last_num := -1
		for pos, pageNumString := range strings.Split(orderingRule, "|") {
			pageNum, err := strconv.Atoi(pageNumString)
			common.Check(err)
			node, ok := rules[pageNum]
			if !ok {
				node = newNode(pageNum)
				rules[pageNum] = node
			}
			if pos == 1 {
				rules[last_num].after = append(rules[last_num].after, node)
				node.before = append(node.before, rules[last_num])
			}
			last_num = pageNum
		}
	}
	for _, pageUpdate := range strings.Split(pageUpdates, "\n") {
		updateNodes := make([]*Node, 0)
		for _, update := range strings.Split(pageUpdate, ",") {
			updatePage, err := strconv.Atoi(update)
			common.Check(err)
			rule, ok := rules[updatePage]
			if !ok {
				e = fmt.Errorf("could not find page %d in rules list", updatePage)
			}
			updateNodes = append(updateNodes, rule)
		}
		updates = append(updates, updateNodes)
	}
	return
}

func updatesValue(updates []*Node) int {
	for i := len(updates); i > 1; i-- {
		if slices.Index(updates[i-2].after, updates[i-1]) < 0 {
			return 0
		}
	}
	return updates[(len(updates)-1)/2].value
}

func Part1(input string) (result int) {
	result = 0
	updates, err := inputParser(input)
	common.Check(err)
	for _, update := range updates {
		result += updatesValue(update)
	}
	logger.Info(fmt.Sprintf("Part 1: %v", result))
	return
}

func correctedResult(updates []*Node) int {
	slices.SortFunc(updates, func(a *Node, b *Node) int {
		if slices.Index(a.after, b) >= 0 {
			return 1
		}
		if slices.Index(b.after, a) >= 0 {
			return -1
		}
		return 0
	})
	return updates[(len(updates)-1)/2].value
}

func Part2(input string) (result int) {
	result = 0
	updates, err := inputParser(input)
	common.Check(err)
	for _, update := range updates {
		if result := updatesValue(update); result > 0 {
			continue
		}
		result += correctedResult(update)
	}
	logger.Info(fmt.Sprintf("Part 2: %v", result))
	return
}

func Day() (part1, part2 int) {
	defer common.Timer("2024-12-05", logger.Info, false)()
	part1 = Part1(input)
	part2 = Part2(input)
	return
}

func main() {
	_, _ = Day()
}
