// **********************************************************************************
// Problem Number: 215
// Problem Name: Kth Largest Element in an Array
// Problem Link: https://leetcode.com/problems/kth-largest-element-in-an-array/
// Status: Solved
// **********************************************************************************

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k := 4
	fmt.Println(findKthLargest(nums, k))
}

func findKthLargest(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	return nums[k-1]
}
