package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"

	common "github.com/KyleJShaver/aoc.go/common"
)

var year, day int = 2024, 07

//go:embed input.txt
var input string
var logger = common.Logger()

const (
	add         = iota
	multiply    = iota
	concatenate = iota
)

func extractEquations(input string) [][]int {
	result := make([][]int, 0)
	for _, equation := range strings.Split(input, "\n") {
		colonIndex := strings.Index(equation, ":")
		if colonIndex < 1 {
			common.Check(fmt.Errorf("could not find a valid colon location, found %v", colonIndex))
		}
		equals, err := strconv.Atoi(equation[:colonIndex])
		common.Check(err)
		operandStrings := strings.Split(equation[colonIndex+2:], " ")
		operands := make([]int, 0, len(operandStrings))
		for _, valString := range operandStrings {
			val, err := strconv.Atoi(valString)
			common.Check(err)
			operands = append(operands, val)
		}
		result = append(result, append([]int{equals}, operands...))
	}
	return result
}

func concatLastTwo(values []int) []int {
	if len(values) < 2 {
		return nil
	}
	left := values[len(values)-2]
	right := values[len(values)-1]
	concatted, err := strconv.Atoi(fmt.Sprintf("%d%d", left, right))
	common.Check(err)
	return append(values[:len(values)-2], []int{concatted}...)
}

func tryOperators(equation []int, operators []int) []string {
	equals := equation[0]
	lastValue := equation[len(equation)-1]
	if equals < 0 {
		return nil
	}
	if len(equation) == 2 {
		if equals == lastValue {
			return []string{""}
		}
		return nil
	}
	possibleOperators := make([]string, 0, len(operators))

	if slices.Contains(operators, multiply) && equals%lastValue == 0 {
		newEquals := equals / lastValue
		possibilities := tryOperators(append([]int{newEquals}, equation[1:len(equation)-1]...), operators)
		for _, possibility := range possibilities {
			possibleOperators = append(possibleOperators, possibility+"*")
		}
	}
	if slices.Contains(operators, add) {
		newEquals := equals - lastValue
		possibilities := tryOperators(append([]int{newEquals}, equation[1:len(equation)-1]...), operators)
		for _, possibility := range possibilities {
			possibleOperators = append(possibleOperators, possibility+"+")
		}
	}
	if slices.Contains(operators, concatenate) {
		equalsString := fmt.Sprintf("%d", equation[0])
		lastString := fmt.Sprintf("%d", equation[len(equation)-1])

		if len(equalsString) > len(lastString)+1 && equalsString[len(equalsString)-len(lastString):] == lastString {
			equalsUnconcatted := equalsString[:len(equalsString)-len(lastString)]
			newEquals, err := strconv.Atoi(equalsUnconcatted)
			common.Check(err)
			possibilities := tryOperators(append([]int{newEquals}, equation[1:len(equation)-1]...), operators)
			for _, possibility := range possibilities {
				possibleOperators = append(possibleOperators, possibility+"|")
			}

		}
	}
	return possibleOperators
}

func printEquation(equation []int, operators string) error {
	equationString := fmt.Sprintf("%v == ", equation[0])
	for range len(equation) - 1 {
		equationString += "("
	}
	operatorsSlice := strings.Split(operators, "")
	total := 0
	for pos := range equation {
		if pos == 0 {
			continue
		}
		if pos > len(operatorsSlice) {
			if equationString[len(equationString)-1] == '+' {
				total += equation[pos]
			} else if equationString[len(equationString)-1] == '*' {
				total *= equation[pos]
			}
			equationString += fmt.Sprintf("%v)", equation[pos])
			continue
		}
		if total == 0 {
			total = equation[pos]
		} else if equationString[len(equationString)-1] == '+' {
			total += equation[pos]
		} else if equationString[len(equationString)-1] == '*' {
			total *= equation[pos]
		}
		equationString += fmt.Sprintf("%v)%s", equation[pos], operatorsSlice[pos-1])

	}
	logger.Debug(equationString)
	if total != equation[0] {
		return fmt.Errorf("expected %v, got %v", equation[0], total)
	}
	return nil
}

func Part1(input string) (result int) {
	result = 0
	equations := extractEquations(input)
	for _, equation := range equations {
		validOperators := tryOperators(equation, []int{add, multiply})
		for pos, operators := range validOperators {
			if pos == 0 {
				result += equation[0]
			}
			err := printEquation(equation, operators)
			common.Check(err)
		}
	}
	logger.Info(fmt.Sprintf("Part 1: %v", result))
	return
}

func Part2(input string) (result int) {
	result = 0
	equations := extractEquations(input)
	for _, equation := range equations {
		validOperators := tryOperators(equation, []int{add, multiply, concatenate})
		for pos := range validOperators {
			if pos == 0 {
				result += equation[0]
			}
		}
	}
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
