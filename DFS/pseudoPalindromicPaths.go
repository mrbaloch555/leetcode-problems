package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pseudoPalindromicPaths(root *TreeNode) int {
	var count int
	pathFreq := make(map[int]int)
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		pathFreq[node.Val]++

		if node.Left == nil && node.Right == nil {
			if isPseudoPalindrome(pathFreq) {
				count++
			}
		} else {
			dfs(node.Left)
			dfs(node.Right)
		}

		pathFreq[node.Val]--
		if pathFreq[node.Val] == 0 {
			delete(pathFreq, node.Val)
		}
	}

	dfs(root)
	return count
}

func isPseudoPalindrome(freq map[int]int) bool {
	var odds int
	for _, val := range freq {
		if val%2 == 1 {
			odds++
		}
		if odds > 1 {
			return false
		}
	}
	return true
}

func main() {
	root := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 1},
		},
		Right: &TreeNode{
			Val:   1,
			Right: &TreeNode{Val: 1},
		},
	}
	fmt.Println(pseudoPalindromicPaths(root))
}
