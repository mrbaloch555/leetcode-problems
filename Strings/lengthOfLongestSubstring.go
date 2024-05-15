package main

import "fmt"

func main() {
	s := "pwwkew"
	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	// Initialize a map to store the last seen index of each character.
	charIndex := make(map[byte]int)
	maxLength := 0
	start := 0 // Start of the sliding window.

	for end := 0; end < len(s); end++ {
		if lastIndex, exists := charIndex[s[end]]; exists {
			fmt.Println(lastIndex)
			// If the character is already in the substring, update the start of the window.
			start = max(start, lastIndex+1)
		}
		// Calculate the length of the current substring.
		substringLength := end - start + 1
		// Update the max length if needed.
		maxLength = max(maxLength, substringLength)
		// Update the last seen index of the current character.
		charIndex[s[end]] = end
		fmt.Println(substringLength, maxLength, start, charIndex)
	}

	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
