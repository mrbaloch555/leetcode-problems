// **********************************************************************************
// Problem Number: 27
// Problem Name: Remove Element
// Problem Link: https://leetcode.com/problems/remove-element/description/
// Status: Solved
// **********************************************************************************

package main

import "fmt"

func main() {
	nums := []int{3, 2, 2, 3}
	val := 3
	fmt.Println(removeElement(nums, val))
}

func removeElement(nums []int, val int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}
	return len(nums)
}
