package common

import (
	"testing"
	"time"
)

func TestReverseString(t *testing.T) {
	inputExpexts := [][]string{
		{"nepalflag🇳🇵", "🇵🇳galflapen"}, // emoji are hard
		{"hello", "olleh"},
		{"", ""},
		{"A", "A"},
	}
	for pos, inputExpext := range inputExpexts {
		got := ReverseString(inputExpext[0])
		expected := inputExpext[1]
		if got != expected {
			t.Errorf("Failed test %d: Expected %s, got %s", pos+1, expected, got)
		}
	}
}

func TestCacher(t *testing.T) {
	inputs := []any{1, "hi"}
	fullSleepTime := time.Millisecond * 200
	sleepThenReturnInput := func(input any) any {
		time.Sleep(fullSleepTime)
		return input
	}
	cachedFn := Cacher(sleepThenReturnInput)
	for pos, input := range inputs {
		start := time.Now()
		got := cachedFn(input)
		uncachcedDuration := time.Since(start)
		if got != input {
			t.Errorf("Failed test %d: Expected %s, got %s from uncached response", pos+1, input, got)
		}
		start = time.Now()
		time.Sleep(fullSleepTime / 10 * 2) // 20% of the uncached time should accoiunt for any wonkiness
		got = cachedFn(input)
		cachcedDuration := time.Since(start)
		if got != input {
			t.Errorf("Failed test %d: Expected %s, got %s from cached response", pos+1, input, got)
		}
		if cachcedDuration >= uncachcedDuration {
			t.Errorf("Cached response took `%v`. This is >= the uncached respnonse time of `%v`", cachcedDuration, uncachcedDuration)
		}
	}
}
