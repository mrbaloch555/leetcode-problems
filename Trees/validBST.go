package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

	root := &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}
	fmt.Println(isValidBST(root))
}

func isValidBST(root *TreeNode) bool {
	
    stack := []TreeNode{*root}

    for len(stack) > 0 {
        current := stack[len(stack) - 1]
        
    }
	return true
}
