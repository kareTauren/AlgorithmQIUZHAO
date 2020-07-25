package home_work

import (
	"fmt"
	"testing"
)

func TestBinaryTreeInOrderTraverse(t *testing.T) {
	node := TNConstruct()
	node.Val = 1
	node.Right = TNNew(2)
	node.Right.Left = TNNew(3)
	fmt.Println(binaryTreeInOrderTraverse(node))
}

func binaryTreeInOrderTraverse(node *TreeNode) []int {
	ans := make([]int, 0)
	if node == nil {
		return ans
	}

	sk := make([]*TreeNode, 0)
	cur := node
	for len(sk) > 0 || cur != nil {
		for cur != nil {
			// push
			sk = append(sk, cur)
			cur = cur.Left
		}
		// pop
		node := sk[len(sk)-1]
		sk = sk[:len(sk)-1]
		ans = append(ans, node.Val)
		cur = node.Right
	}
	return ans
}
