package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var result strings.Builder

	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			result.WriteString("nil,")
			return
		}
		result.WriteString(strconv.Itoa(node.Val))
		result.WriteByte(',')
		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)
	return result.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	values := strings.Split(data, ",")
	return this.deserializeHelper(&values)
}

func (this *Codec) deserializeHelper(data *[]string) *TreeNode {
	if len(*data) == 0 || (*data)[0] == "nil" {
		*data = (*data)[1:]
		return nil
	}
	val, _ := strconv.Atoi((*data)[0])
	root := &TreeNode{Val: val}
	*data = (*data)[1:]
	root.Left = this.deserializeHelper(data)
	root.Right = this.deserializeHelper(data)
	return root
}

func main() {
	root := &TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 1, Left: &TreeNode{Val: 4}},
		Right: &TreeNode{Val: 3},
	}
	obj := Constructor()
	str := obj.serialize(root)
	root = obj.deserialize(str)
	fmt.Println(root.Left, root.Right, root.Val)
}
