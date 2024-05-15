package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	Tree     *TreeNode
	Elements []int
}

func Constructor(root *TreeNode) BSTIterator {
	elements := inorderTraversal(root)
	return BSTIterator{
		Tree:     root,
		Elements: elements,
	}
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	leftNodes := inorderTraversal(root.Left)
	rightNodes := inorderTraversal(root.Right)
	return append(leftNodes, append([]int{root.Val}, rightNodes...)...)
}

func (this *BSTIterator) Next() int {
	current := this.Elements[0]
	this.Elements = this.Elements[1:]
	return current
}

func (this *BSTIterator) HasNext() bool {
	return len(this.Elements) != 0
}

func main() {
	root := &TreeNode{Val: 7, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 15, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20}}}
	obj := Constructor(root)
	fmt.Println(obj.Next())
	fmt.Println(obj.Next())
	fmt.Println(obj.HasNext())
	fmt.Println(obj.Next())
	fmt.Println(obj.HasNext())
	fmt.Println(obj.Next())
	fmt.Println(obj.HasNext())
	fmt.Println(obj.Next())
	fmt.Println(obj.HasNext())
}
