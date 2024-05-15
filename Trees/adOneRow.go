package main

import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
    if depth == 1 {
        newRoot := &TreeNode{Val: val}
        newRoot.Left = root;
        return newRoot;
    }
    queue := []*TreeNode{root}
	for level := 1; level < depth-1; level++ {
		nextLevel := []*TreeNode{}
		for _, node := range queue {
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		queue = nextLevel
	}
    for i := 0; i < len(queue); i++ {
        left := &TreeNode{Val: val, Left: queue[i].Left}
        right := &TreeNode{Val: val, Right: queue[i].Right}
        queue[i].Left = left;
        queue[i].Right = right
    }
    return root
}

func main(){
    root := &TreeNode{
        Val: 4,
        Left: &TreeNode{
            Val: 2,
            Left: &TreeNode{Val: 3},
            Right: &TreeNode{Val: 1},
        },
        Right: &TreeNode{
            Val: 6,
            Left: &TreeNode{Val: 5},
        },
    }
    root = addOneRow(root, 1, 2);
    stack := []*TreeNode{root}
    for len(stack) > 0 {
        curr := stack[0]
        stack = stack[1:]
        if curr.Left != nil {stack = append(stack, curr.Left)}
        if curr.Right != nil {stack = append(stack, curr.Right)}
        fmt.Println(curr.Val)
    }
}
