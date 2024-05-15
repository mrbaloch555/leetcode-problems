package main

import (
	"fmt"
	"math"
)

func minSubArrayLen(target int, nums []int) int {
	minLen := math.MaxInt32
	currSum, l := 0, 0

	for i, val := range nums {
		currSum += val

		for currSum >= target {
			minLen = min(i-l+1, minLen)
			currSum -= nums[l]
			l += 1
		}
	}
	if minLen == math.MaxInt32 {
		return 0
	} else {
		return minLen
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {

	target := 7
	nums := []int{2, 3, 1, 2, 4, 3}
	fmt.Println(minSubArrayLen(target, nums))
}
