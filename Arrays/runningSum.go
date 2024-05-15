package main

// **********************************************************************************
// Problem Number: 1480
// Problem Name: Running Sum of 1d Array
// Problem Link: https://leetcode.com/problems/running-sum-of-1d-array/description/
// Status: Solved
// **********************************************************************************

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	nums = runningSum(nums)
	fmt.Println(nums)
}
func runningSum(nums []int) []int {
	sum := []int{}
	counter := 0
	for i, _ := range nums {
		counter += nums[i]
		sum = append(sum, counter)
	}
	return sum
}
