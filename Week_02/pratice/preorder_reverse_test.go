package pratice

import (
	"fmt"
	"testing"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func construct() *TreeNode {
	node := New(1)
	node.Right = New(4)
	node.Left = New(3)
	node.Left.Left = New(5)
	node.Left.Right = New(6)
	return node
}

func New(val int) *TreeNode {
	return &TreeNode{Val: val}
}

func TestPreorderTraversal(t *testing.T) {
	node := construct()
	fmt.Print(preorderTraversal(node))
}

// 前序遍历 根 - 左 - 右 方法一
// func preorderTraversal(root *TreeNode) []int {
// 	ans := []int{}
// 	stack := stack{root}
//
// 	for len(stack) != 0 {
// 		node := stack.Pop()
// 		ans = append(ans, node.Val)
// 		if node.Left != nil {
// 			stack.Push(node.Left)
// 		}
// 		if node.Right != nil {
// 			stack.Push(node.Right)
// 		}
// 	}
// 	return ans
// }

// 方法二
func preorderTraversal(root *TreeNode) []int {
	ans := []int{}
	stack := stack{}
	cur := root

	for cur != nil || len(stack) > 0 {
		for cur != nil {
			ans = append(ans, cur.Val)
			stack.Push(cur)
			cur = cur.Left
		}
		tmp := stack.Pop()
		cur = tmp.Right
	}
	return ans
}
