package main

import (
	"fmt"
	"math"
)

func maximumValueSum(nums []int, k int, edges [][]int) int64 {
	maxSum := 0
	changedCount := 0

	for _, num := range nums {
		maxSum += max(num, num^k)
		if (num ^ k) > num {
			changedCount++
		}
	}

	if changedCount%2 == 0 {
		return int64(maxSum)
	}

	minChangeDiff := math.MaxInt64
	for _, num := range nums {
		changeDiff := abs(num - (num ^ k))
		if changeDiff < minChangeDiff {
			minChangeDiff = changeDiff
		}
	}

	return int64(maxSum - minChangeDiff)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func main() {
	nums := []int{7, 7, 7, 7, 7, 7}
	k := 3
	edges := [][]int{{0, 1}, {0, 2}, {0, 3}, {0, 4}, {0, 5}}
	fmt.Println(maximumValueSum(nums, k, edges))
}
