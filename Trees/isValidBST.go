package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return helper(root, nil, nil)
}

func helper(root *TreeNode, min *int, max *int) bool {
	if root == nil {
		return true
	}
	if min != nil && root.Val <= *min || max != nil && root.Val >= *max {
		return false
	}
	return helper(root.Left, min, &root.Val) && helper(root.Right, &root.Val, max)

}
func main() {
	root := &TreeNode{Val: 5, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 1, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 6}}}
	fmt.Println(isValidBST(root))
}
