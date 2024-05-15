package main

import "fmt"

func maxSubarraySumCircular(nums []int) int {
	currSum := nums[len(nums) -1];
	maxSum := nums[len(nums) -1]
    l := len(nums) - 1;
    fmt.Println(currSum)
	for i := 0; i < len(nums); i++ {
		currSum = max(currSum, nums[i])
        fmt.Println(currSum)
        if i != l {
		    currSum += nums[i]
        }
		maxSum = max(currSum, maxSum)
	}
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	nums := []int{-3, -2, -3}
	fmt.Println(maxSubarraySumCircular(nums))
}
