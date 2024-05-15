package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	Val  int
	Next *Node
}

func main() {
	var a = Node{Val: 2}
	var b = Node{Val: 4}
	var c = Node{Val: 3}
	var d = Node{Val: 5}
	var e = Node{Val: 6}
	var f = Node{Val: 4}

	a.Next = &b
	b.Next = &c
	d.Next = &e
	e.Next = &f

	var head Node
	addTwoNumbers(head, &a, &d)
	fmt.Println(head)
	// for head != nil {
	// 	fmt.Println(head.Val)
	// 	head = head.Next
	// }
	// printLikedList(&a)
}

func addTwoNumbers(head Node, l1 *Node, l2 *Node) *Node {
	l1_current := l1
	l2_current := l2
	l1_sum := ""
	l2_sum := ""
	for l1_current != nil {
		l1_sum = fmt.Sprintf("%s%d", l1_sum, l1_current.Val)
		l1_current = l1_current.Next
	}
	for l2_current != nil {
		l2_sum = fmt.Sprintf("%s%d", l2_sum, l2_current.Val)
		l2_current = l2_current.Next
	}
	l1_sum_int, _ := strconv.ParseInt(l1_sum, 10, 0)
	l2_sum_int, _ := strconv.ParseInt(l2_sum, 10, 0)

	total_sum := strconv.Itoa(int(l1_sum_int) + int(l2_sum_int))

	nodes := strings.Split(total_sum, "")
	add(&head, nodes, 0)
	return &head
}

func add(head *Node, nodes []string, counter int) {
	if counter >= len(nodes) {
		return
	}
	// fmt.Println(nodes[counter], counter)
	val, _ := strconv.ParseInt(nodes[counter], 10, 0)
	head.Next = &Node{Val: int(val)}
	counter++
	add(head.Next, nodes, counter)
}
