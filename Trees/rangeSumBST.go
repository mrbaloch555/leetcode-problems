// **********************************************************************************
// Problem Number: 938
// Problem Name: Range Sum of BST
// Problem Link: https://leetcode.com/problems/range-sum-of-bst/
// Status: Solved
// **********************************************************************************

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	a := TreeNode{Val: 10}
	b := TreeNode{Val: 5}
	c := TreeNode{Val: 3}
	d := TreeNode{Val: 7}
	e := TreeNode{Val: 15}
	f := TreeNode{Val: 18}

	a.Left = &b
	a.Right = &e
	b.Left = &c
	b.Right = &d
	e.Right = &f

	low := 7
	high := 15
	fmt.Println(rangeSumBST(&a, low, high))
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	sum := 0
	stack := []*TreeNode{root}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if current.Left != nil {
			stack = append(stack, current.Left)
		}
		if current.Right != nil {
			stack = append(stack, current.Right)
		}

		if current.Val >= low && current.Val <= high {
			sum += current.Val
		}
	}
	return sum
}
