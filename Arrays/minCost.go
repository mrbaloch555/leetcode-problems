package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{1, 3, 5, 2}
	cost := []int{2, 3, 1, 14}
	minCost := minCost(nums, cost)
	fmt.Println("Minimum cost to make the array equal:", minCost)
}

func minCost(nums []int, cost []int) int {
	maxNum := math.MinInt32
	totalCost := 0

	// Find the maximum number in the array
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
	}

	// Calculate the total cost
	for i := 0; i < len(nums); i++ {
		if nums[i] != maxNum {
			totalCost += cost[i]
		}
	}

	return totalCost
}
