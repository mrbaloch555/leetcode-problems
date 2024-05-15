// **********************************************************************************
// Problem Number: 2460
// Problem Name: Apply Operations to an Array
// Problem Link: https://leetcode.com/problems/apply-operations-to-an-array/
// Status: Solved
// **********************************************************************************
package main

import "fmt"

func main() {
	nums := []int{0, 1}
	fmt.Println(applyOperations(nums))
}

func applyOperations(nums []int) []int {
	zero_counts := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums = append(nums[:i], nums[i+1:]...)
			i--
			zero_counts = append(zero_counts, 0)
		} else if i < len(nums)-1 && nums[i] == nums[i+1] && nums[i] != 0 {
			nums[i] = nums[i] * 2
			nums = append(nums[:i+1], nums[i+2:]...)
			zero_counts = append(zero_counts, 0)
		}
	}
	return append(nums, zero_counts...)
}
