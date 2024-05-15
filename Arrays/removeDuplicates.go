// **********************************************************************************
// Problem Number: 26
// Problem Name: Remove Duplicates from Sorted Array
// Problem Link: https://leetcode.com/problems/remove-duplicates-from-sorted-array/
// Status: Solved
// **********************************************************************************
package main

import "fmt"

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 3, 4, 4}
	fmt.Println(removeDuplicates(nums))
}

func removeDuplicates(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if i < len(nums)-1 && nums[i] == nums[i+1] {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}

	return len(nums)
}
