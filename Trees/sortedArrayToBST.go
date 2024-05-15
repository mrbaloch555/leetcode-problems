// **********************************************************************************
// Problem Number: 108
// Problem Name: Convert Sorted Array to Binary Search Tree
// Problem Link: https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/
// Status: Solved
// **********************************************************************************

package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	nums := []int{-10, -3, 0, 5, 9}
	fmt.Println(sortedArrayToBST(nums))
}
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := int(math.Floor(float64(len(nums) / 2)))
	root := &TreeNode{Val: nums[mid]}

	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])

	return root
}
