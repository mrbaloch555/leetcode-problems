package main

import (
	"fmt"
)

// func lengthOfLongestSubstring(s string) int {

// 	if len(s) == 0 {
// 		return 0
// 	}
// 	maxLen := 0
// 	left := 0
// 	for i := 0; i < len(s); i++ {
// 		if left != i && strings.Contains(s[left:i], string(s[i])) {
// 			sub := s[left:i]
// 			if len(sub) > maxLen {
// 				maxLen = len(sub)
// 			}
// 			left++
// 			i = left
// 		}
// 		if i == len(s)-1 {
// 			sub := s[left : i+1]
// 			if len(sub) > maxLen {
// 				maxLen = len(sub)
// 			}
// 		}
// 	}

// 	return maxLen
// }

// func lengthOfLongestSubstring(s string) int {
// 	charIndex := make(map[byte]int)
// 	maxLen := 0
// 	start := 0

// 	for end := 0; end < len(s); end++ {
// 		if index, found := charIndex[s[end]]; found && index >= start {
// 			start = index + 1
// 		}

// 		if curLen := end - start + 1; curLen > maxLen {
// 			maxLen = curLen
// 		}

// 		charIndex[s[end]] = end
// 	}

// 	return maxLen
// }

func lengthOfLongestSubstring(s string) int {
	charMap := map[rune]int{}
	maxLen := 0
	l := 0
	for i, val := range s {
		pos, ok := charMap[val]
		if ok {
			l = max(l, pos+1)
		}
		charMap[val] = i
		maxLen = max(maxLen, i-l+1)
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	s := "pwwkew"
	fmt.Println(lengthOfLongestSubstring(s))
}
