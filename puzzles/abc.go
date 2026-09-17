package main

import (
	"fmt"
	"strings"
)

func canMakeWord(word string, blocks []string) bool {
	used := make([]bool, len(blocks)) // all false
	return canMakeWordHelper(word, blocks, used, 0)
}

func canMakeWordHelper(word string, blocks []string, used []bool, idx int) bool {
	if idx == len(word) { // every letter has been matched to a block
		return true
	}

	c := word[idx] // current letter
	for i, blk := range blocks {
		if used[i] || !strings.ContainsRune(blk, rune(c)) { // already checked, or doesn't contain the letter
			continue
		}

		// recurse
		used[i] = true
		if canMakeWordHelper(word, blocks, used, idx+1) {
			return true
		}
		used[i] = false
	}

	return false
}

func main() {
	blocks := []string{
		"BO", "XK", "DQ", "CP", "NA", "GT", "RE",
		"TG", "QD", "FS", "JW", "HU", "VI", "AN",
		"OB", "ER", "FS", "LY", "PC", "ZM",
	}
	words := []string{"A", "BARK", "BOOK", "TREAT", "COMMON", "SQUAD", "CONFUSE"}

	for _, w := range words {
		fmt.Println("Can make", w, ":", canMakeWord(w, blocks))
	}
}
