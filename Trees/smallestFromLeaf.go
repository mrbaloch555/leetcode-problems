package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var charMap = map[int]byte{
	0:  'a',
	1:  'b',
	2:  'c',
	3:  'd',
	4:  'e',
	5:  'f',
	6:  'g',
	7:  'h',
	8:  'i',
	9:  'j',
	10: 'k',
	11: 'l',
	12: 'm',
	13: 'n',
	14: 'o',
	15: 'p',
	16: 'q',
	17: 'r',
	18: 's',
	19: 't',
	20: 'u',
	21: 'v',
	22: 'w',
	23: 'x',
	24: 'y',
	25: 'z',
}

func smallestFromLeaf(root *TreeNode) string {
	if root == nil {
		return ""
	}
	return dfs(root, "")
}

func dfs(node *TreeNode, suffix string) string {
	if node == nil {
		return ""
	}

	currentStr := string(charMap[node.Val]) + suffix

	if node.Left == nil && node.Right == nil {
		return currentStr
	}

	leftStr := dfs(node.Left, currentStr)
	rightStr := dfs(node.Right, currentStr)

	if leftStr == "" {
		return rightStr
	} else if rightStr == "" {
		return leftStr
	} else {
		if leftStr < rightStr {
			return leftStr
		} else {
			return rightStr
		}
	}
}

func main() {
	root := &TreeNode{
		Val:   0,
		Left:  &TreeNode{Val: 1, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}},
		Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}},
	}

	fmt.Println(smallestFromLeaf(root))
}
