package main

import (
	"fmt"
)

func maxSubArray(nums []int) int {
	currSum := 0
	maxSum := nums[0]

	for _, val := range nums {
		currSum = max(currSum, 0)
		currSum += val
		maxSum = max(currSum, maxSum)
	}
	return maxSum
}

func maxSubArrayLength(nums []int) []int {
	currSum := 0
	maxL, maxR := 0, 0
	maxSum := nums[0]
	L := 0

	for i, val := range nums {
		if currSum < 0 {
			currSum = 0
			L = i
		}
		currSum += val
		if currSum > maxSum {
			maxSum = currSum
			maxL, maxR = L, i
		}
	}
	return []int{maxL, maxR}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	nums := []int{4, -1, 2, -7, 3, 4}
	fmt.Println(maxSubArray(nums))
	fmt.Println(maxSubArrayLength(nums))
}
