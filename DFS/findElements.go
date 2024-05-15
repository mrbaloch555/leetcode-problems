package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type FindElements struct {
	elements map[int]int
}

func Constructor(root *TreeNode) FindElements {
	root.Val = 0
	elements := map[int]int{root.Val: root.Val}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root.Left != nil {
			root.Left.Val = (2 * root.Val) + 1
			elements[root.Left.Val] = root.Left.Val
			dfs(root.Left)
		}
		if root.Right != nil {
			root.Right.Val = (2 * root.Val) + 2
			elements[root.Right.Val] = root.Right.Val
			dfs(root.Right)
		}
	}
	dfs(root)
	return FindElements{
		elements: elements,
	}
}

func (this *FindElements) Find(target int) bool {
	if _, ok := this.elements[target]; ok {
		return true
	}
	return false
}

func main() {
	root := &TreeNode{
		Val: -1,
		Left: &TreeNode{
			Val:   -1,
			Left:  &TreeNode{Val: -1},
			Right: &TreeNode{Val: -1},
		},
		Right: &TreeNode{Val: -1},
	}
	obj := Constructor(root)
	fmt.Println(obj.Find(1))
	fmt.Println(obj.Find(3))
	fmt.Println(obj.Find(5))
}
