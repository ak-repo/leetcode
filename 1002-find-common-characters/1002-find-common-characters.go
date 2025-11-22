package main

import (
	"fmt"
)

func commonChars(words []string) []string {
	if len(words) == 0 {
		return []string{}
	}

	minFreq := make(map[rune]int)
	for _, ch := range words[0] {
		minFreq[ch]++
	}

	for _, word := range words[1:] {
		currFreq := make(map[rune]int)
		for _, ch := range word {
			currFreq[ch]++
		}
		for ch := range minFreq {
			if currFreq[ch] < minFreq[ch] {
				minFreq[ch] = currFreq[ch]
			}
		}
	}

	result := []string{}
	for ch, count := range minFreq {
		for i := 0; i < count; i++ {
			result = append(result, string(ch))
		}
	}

	return result
}


