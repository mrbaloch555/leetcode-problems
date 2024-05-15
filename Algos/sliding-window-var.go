package main

import (
	"fmt"
	"math"
)

func longestSubArray(nums []int) int {
	longest := 0
	l := 0

	for i, val := range nums {
		if val != nums[l] {
			l = i
		}
		longest = max(longest, i-l+1)
	}
	return longest
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func minimumLength(nums []int, target int) int {
	minLen := math.MaxInt32
	l := 0
	currSum := 0
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

func longestSubArrayNavie(nums []int) []int {
	longest := []int{}
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			sub := nums[i : j+1]
			if every(sub, func(prev, curr int) bool {
				return prev == curr
			}) {
				if len(longest) < len(sub) {
					longest = sub
				}
			}
		}
	}
	return longest
}

func every(nums []int, filter func(prev, curr int) bool) bool {
	isSatisfied := true
	prev := nums[0]
	for _, val := range nums[1:] {
		if !filter(prev, val) {
			isSatisfied = false
			return isSatisfied
		}
	}
	return isSatisfied
}

func main() {
	nums := []int{2, 3, 1, 2, 4, 3}
	fmt.Println(longestSubArrayNavie(nums))
	fmt.Println(longestSubArray(nums))
	fmt.Println(minimumLength(nums, 7))
}
