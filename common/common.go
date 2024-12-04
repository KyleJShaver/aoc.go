package common

import (
	"cmp"
	"fmt"
	"log/slog"
	"os"
	"sort"
	"strconv"
	"strings"
	"testing"
	"time"
)

// `if err != nil {panic(err)}`
func Check(err error) {
	if err != nil {
		panic(err)
	}
}

// Sorts the provided slice in-place, as long as the type is in [cmp.Ordered].
func SortSlice[C cmp.Ordered](list []C) {
	sort.SliceStable(list, func(i, j int) bool {
		return list[i] < list[j]
	})
}

// Configures a default JSON Structured Logger to stdout
func Logger() *slog.Logger {
	slog.SetLogLoggerLevel(slog.LevelDebug)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	slog.SetDefault(logger)
	return logger
}

// Creates a timer closure. Best used with [defer], like so:
//
//	defer common.Timer("2024-12-01", logger.Info, false)()
//
// This will log that the timer has started, and will log the elapsed time at the end of the function.
func Timer(name string, fn func(msg string, args ...any), newline bool) func() {
	start := time.Now()
	nlChar := ""
	if newline {
		nlChar = "\n"
	}
	fn(fmt.Sprintf("Starting %s...%s", name, nlChar))
	return func() {
		fn(fmt.Sprintf("Finished %s in %v%s", name, time.Since(start), nlChar))
	}
}

// Creates a closure where the returned function's responses are cached
func CachedListInts(split string) func(input string) [][]int {
	cache := make(map[string][][]int)
	return func(input string) (nums [][]int) {
		cacheKey := input + split
		if retVal, ok := cache[cacheKey]; ok {
			return retVal
		}
		lines := strings.Split(input, "\n")
		nums = make([][]int, 0, len(lines))
		for _, line := range lines {
			components := strings.Split(line, split)
			levels := make([]int, 0, len(components))
			for _, component := range components {
				intVal, err := strconv.Atoi(component)
				Check(err)
				levels = append(levels, intVal)
			}
			nums = append(nums, levels)
		}
		cache[cacheKey] = nums
		return
	}
}

type AOCTest struct {
	Expected int
	Input    string
	Fn       func(string) int
	Label    string
}

func (tst AOCTest) Run(t *testing.T) {
	t.Run(tst.Label, func(t *testing.T) {
		if got := tst.Fn(tst.Input); got != tst.Expected {
			t.Errorf("%s: Expected `%v`, got `%v`", tst.Label, tst.Expected, got)
		}
	})
}

func Grid(input string) (grid [][]string) {
	grid = make([][]string, 0)
	for _, row := range strings.Split(input, "\n") {
		grid = append(grid, strings.Split(row, ""))
	}
	return
}

var GridCached = InputCacher(Grid)

func InputCacher[IT comparable, RT any](fn func(IT) RT) func(IT) RT {
	cache := make(map[IT]RT)
	return func(input IT) RT {
		if cached, ok := cache[input]; ok {
			return cached
		}
		cache[input] = fn(input)
		return cache[input]
	}
}

func ReverseString(s string) string {
	r := make([]rune, 0, len(s))
	sRunes := []rune(s)
	for i := len(sRunes); i > 0; i-- {
		r = append(r, sRunes[i-1])
	}
	return string(r)
}
