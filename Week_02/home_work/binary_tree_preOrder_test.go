package home_work

import (
	"fmt"
	"testing"
)

func TestBinaryTreePreOrderTraverse(t *testing.T) {
	node := TNConstruct()
	node.Val = 1
	node.Right = TNNew(2)
	node.Right.Left = TNNew(3)
	fmt.Println(binaryTreePreOrderTraverse(node))
}

func binaryTreePreOrderTraverse(node *TreeNode) []int {
	ans := make([]int, 0)
	if node == nil {
		return ans
	}

	stack := make([]*TreeNode, 0)
	cur := node
	for len(stack) > 0 || cur != nil {
		for cur != nil {
			ans = append(ans, cur.Val)
			stack = append(stack, cur)
			cur = cur.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		cur = node.Right
	}
	return ans
}
