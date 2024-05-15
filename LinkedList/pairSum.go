// **********************************************************************************
// Problem Number: 2130
// Problem Name: Maximum Twin Sum of a Linked List
// Problem Link: https://leetcode.com/problems/maximum-twin-sum-of-a-linked-list/
// Status: Solved
// **********************************************************************************

package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func main() {
	var a = Node{Val: 1}
	var b = Node{Val: 1000}
	var c = Node{Val: 2}
	var d = Node{Val: 3}

	a.Next = &b
	b.Next = &c
	c.Next = &d
	fmt.Println(pairSum(&a))
}

func pairSum(head *Node) int {
	var array = make([]int, 0)
	current := head
	for current != nil {
		array = append(array, current.Val)
		current = current.Next
	}
	var max = array[0]

	if len(array) > 0 {
		for i := 0; i < len(array)/2+1; i++ {
			twin := len(array) - 1 - i
			if array[i]+array[twin] > max {
				max = array[i] + array[twin]
			}
		}
	}
	return max
}
