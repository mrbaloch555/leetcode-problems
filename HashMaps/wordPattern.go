package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	words := strings.Fields(s)

	if len(pattern) != len(words) {
		return false
	}

	patternToWord := make(map[byte]string)
	wordToPattern := make(map[string]byte)

	for i := 0; i < len(pattern); i++ {
		p := pattern[i]
		word := words[i]

		if prevWord, exists := patternToWord[p]; exists {
			if prevWord != word {
				return false
			}
		} else {
			patternToWord[p] = word
		}

		if prevPattern, exists := wordToPattern[word]; exists {
			if prevPattern != p {
				return false
			}
		} else {
			wordToPattern[word] = p
		}
	}

	return true
}

func main() {
	fmt.Println(wordPattern("abba", "dog cat cat dog"))  // Output: true
	fmt.Println(wordPattern("abba", "dog cat cat fish")) // Output: false
	fmt.Println(wordPattern("aaaa", "dog cat cat dog"))  // Output: false
}
