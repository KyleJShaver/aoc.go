package common

import "testing"

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
