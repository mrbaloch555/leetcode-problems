package main

import "fmt"

// Find a non empty subarray with the largest sum
func largestSumNavie(nums []int) int {
	max := nums[0]
	for i, _ := range nums {
		sum := 0
		for _, inner := range nums[i:] {
			sum += inner
		}
		if sum > max {
			max = sum
		}
	}
	return max
}

func kadanes(nums []int) int {
	maxSum := nums[0]
	currSum := 0

	for _, val := range nums {
		currSum = max(currSum, 0)
		currSum += val
		maxSum = max(maxSum, currSum)
	}
	return maxSum
}

func slidingWindow(nums []int) []int {
	maxSum := nums[0]
	currSum := 0
	maxL, maxR := 0, 0
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
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(largestSumNavie(nums))
	fmt.Println(kadanes(nums))
	fmt.Println(slidingWindow(nums))
}
