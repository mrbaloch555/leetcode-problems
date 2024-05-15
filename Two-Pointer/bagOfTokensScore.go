package main

import (
	"fmt"
	"sort"
)

func bagOfTokensScore(tokens []int, power int) int {
	sort.Ints(tokens)
	left, right := 0, len(tokens)-1
	score := 0
	max_score := 0

	for left <= right {
		if power >= tokens[left] {
			power -= tokens[left]
			score += 1
			left += 1
			max_score = max(max_score, score)
		} else if score > 0 {
			power += tokens[right]
			score -= 1
			right -= 1
		} else {
			break
		}
	}
	return max_score
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	tokens := []int{100, 200, 300, 400}
	power := 200
	fmt.Println(bagOfTokensScore(tokens, power))
}
