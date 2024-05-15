package main

import (
	"fmt"
	"sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{}
	l2 := &ListNode{}

	fmt.Println(mergeTwoLists(l1, l2))

}
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	l1_values := traverse(list1)
	l2_values := traverse(list2)
	merged_array := append(l1_values, l2_values...)
	sort.Slice(merged_array, func(i, j int) bool {
		return merged_array[i] < merged_array[j]
	})
	if len(merged_array) > 0 {
		head := &ListNode{Val: merged_array[0]}
		tail := head
		for _, val := range merged_array[1:] {
			node := &ListNode{Val: val}
			tail.Next = node
			tail = node
		}
		return head
	} else {
		return list1
	}
}

func traverse(head *ListNode) []int {
	values := []int{}

	current := head
	for current != nil {
		values = append(values, current.Val)
		current = current.Next
	}
	return values
}

func add(head *ListNode, val int) *ListNode {
	head = &ListNode{Val: val, Next: &ListNode{}}
	return head.Next
}
