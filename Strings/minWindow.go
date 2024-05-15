// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func minWindow(s string, t string) string {
// 	min := ""
// 	for i := len(s); i >= 0; i-- {
// 		for j := 0; j < len(s)-i; j++ {
// 			sub := s[j : i+j+1]
// 			tmpSub := sub
// 			matched := true
// 			if len(sub) >= len(t) {
// 				for k := 0; k < len(t); k++ {
// 					index := strings.IndexAny(tmpSub, string(t[k]))
// 					if index == -1 {
// 						matched = false
// 						break
// 					} else {
// 						tmpSub = tmpSub[:index] + tmpSub[index+1:]
// 					}
// 				}
// 				if matched {
// 					if min == "" {
// 						min = sub
// 					} else {
// 						if len(sub) < len(min) {
// 							min = sub
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return min
// }

//	func main() {
//		s := "abc"
//		t := "ac"
//		fmt.Println(minWindow(s, t))
//	}
package main

import (
	"fmt"
	"math"
)

func minWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}

	targetFreq := make(map[byte]int)
	for i := range t {
		targetFreq[t[i]]++
	}

	left, right := 0, 0
	minLen := math.MaxInt32
	minWindow := ""

	required := len(t)

	for right < len(s) {
		if targetFreq[s[right]] > 0 {
			required--
		}
		targetFreq[s[right]]--

		for required == 0 {
			if right-left+1 < minLen {
				minLen = right - left + 1
				minWindow = s[left : right+1]
			}

			targetFreq[s[left]]++
			if targetFreq[s[left]] > 0 {
				required++
			}
			left++
		}

		right++
	}

	return minWindow
}

func main() {
	s := "a"
	t := "aa"
	fmt.Println(minWindow(s, t))
}
