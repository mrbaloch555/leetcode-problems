package main

import "fmt"

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []*Node{root}

	res := [][]int{{root.Val}}
	for len(queue) > 0 {
		nextLevel := []*Node{}
		for len(queue) > 0 {
			curr := queue[0]
			queue = queue[1:]

			if len(curr.Children) > 0 {
				nextLevel = append(nextLevel, curr.Children...)
			}
		}
		r := []int{}
		for _, val := range nextLevel {
			r = append(r, val.Val)
		}
		if len(r) > 0 {
			res = append(res, r)
		}
		queue = append(queue, nextLevel...)
	}
	return res
}

func main() {

	root := &Node{
		Val: 1,
		Children: []*Node{
			{Val: 3, Children: []*Node{
				{Val: 5},
				{Val: 6},
			}},
			{Val: 2},
			{Val: 4},
		},
	}
	fmt.Println(levelOrder(root))
}
