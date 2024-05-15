package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var a = ListNode{Val: 7}
	var b = ListNode{Val: 9}
	var c = ListNode{Val: 6}
	var d = ListNode{Val: 6}
	var e = ListNode{Val: 7}
	var f = ListNode{Val: 8}
	var g = ListNode{Val: 3}
	var h = ListNode{Val: 0}
	var i = ListNode{Val: 9}
	var j = ListNode{Val: 5}

	a.Next = &b
	b.Next = &c
	c.Next = &d
	d.Next = &e
	e.Next = &f
	f.Next = &g
	g.Next = &h
	h.Next = &i
	i.Next = &j
	k := 5
	head := swapListNodes(&a, k)

	current := head
	for current != nil {
		fmt.Println(current)
		current = current.Next
	}
}

func swapListNodes(head *ListNode, k int) *ListNode {
	// create counter for reverse count when k becomes 1 we start swaping
	counter := k
	current := head
	val := 0
	for current != nil {
		if k == 1 && counter != 1 {
			// do swaping

			current = current.Next
			val = current.Val
			// make counter decrement so we can end at specific time
		} else {
			k--
			if k == 1 {
				fmt.Println(current.Next.Val)
			}
			// for reverse counting
			counter++
			current = current.Next
		}
	}
	fmt.Println(val)
	return head
}
