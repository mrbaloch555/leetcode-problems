// **********************************************************************************
// Problem Number: 268
// Problem Name: Missing Number
// Problem Link: https://leetcode.com/problems/missing-number/
// Status: Un Solved
// **********************************************************************************

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{0, 1}
	fmt.Println(missingNumber(nums))
}

func missingNumber(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for i := 0; i < len(nums); i++ {
		if len(nums[i:]) >= 2 {
			if nums[i+1]-nums[i] >= 2 {
				return nums[i] + 1
			}
		}
	}
	return nums[0] - 1
}
