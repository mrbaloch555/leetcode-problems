// **********************************************************************************
// Problem Number: 2215
// Problem Name: Find the Difference of Two Arrays
// Problem Link: https://leetcode.com/problems/find-the-difference-of-two-arrays/
// Status: Solved
// **********************************************************************************

package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3}
	nums2 := []int{2, 4, 6}

	fmt.Println(findDifference(nums1, nums2))
}

func findDifference(nums1 []int, nums2 []int) [][]int {
	return func() [][]int {
		obj := make(map[int]bool)
		result := [][]int{{}, {}}
		for _, val := range nums1 {
			found := false
			for _, innerVal := range nums2 {
				if val == innerVal {
					found = true
				}
			}
			if !found && !obj[val] {
				obj[val] = true
				result[0] = append(result[0], val)
			}
		}
		for _, val := range nums2 {
			found := false
			for _, innerVal := range nums1 {
				if val == innerVal {
					found = true
				}
			}
			if !found && !obj[val] {
				obj[val] = true
				result[1] = append(result[1], val)
			}
		}
		return result
	}()
}
