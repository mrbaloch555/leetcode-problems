package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isEvenOddTree(root *TreeNode) bool {
	if root == nil {
		return false
	}
	queue := []*TreeNode{root}
	level := 0

	for len(queue) > 0 {
		isEven := level%2 == 0
		index := -1
		nextLevel := []*TreeNode{}

		for _, current := range queue {
			if isEven && current.Val%2 == 0 {
				return false
			} else if !isEven && current.Val%2 != 0 {
				return false
			}

			if index != -1 {
				if (isEven && index >= current.Val) || (!isEven && index <= current.Val) {
					return false
				}
			}

			index = current.Val

			if current.Left != nil {
				nextLevel = append(nextLevel, current.Left)
			}
			if current.Right != nil {
				nextLevel = append(nextLevel, current.Right)
			}
		}

		level++
		queue = nextLevel
	}
	return true
}

func main() {
	// root := &TreeNode{
	// 	Val: 1,
	// 	Left: &TreeNode{
	// 		Val: 10,
	// 		Left: &TreeNode{
	// 			Val: 3,
	// 			Left: &TreeNode{
	// 				Val: 12,
	// 			},
	// 			Right: &TreeNode{
	// 				Val: 8,
	// 			},
	// 		},
	// 	},
	// 	Right: &TreeNode{
	// 		Val: 4,
	// 		Left: &TreeNode{
	// 			Val: 7,
	// 			Left: &TreeNode{
	// 				Val: 6,
	// 			},
	// 		},
	// 		Right: &TreeNode{
	// 			Val: 9,
	// 			Right: &TreeNode{
	// 				Val: 2,
	// 			},
	// 		},
	// 	},
	// }
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 9,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	fmt.Println(isEvenOddTree(root))
}
