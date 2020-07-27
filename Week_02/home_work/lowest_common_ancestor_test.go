package home_work

import (
	"fmt"
	"testing"
)

func TestLowestCommonAncestor(t *testing.T) {
	p := &TreeNode{Val: 6}
	q := &TreeNode{Val: 2}

	root := TNConstruct()
	root.Val = 3

	root.Left = TNNew(5)
	root.Left.Left = TNNew(6)
	root.Left.Right = TNNew(2)
	root.Left.Right.Left = TNNew(7)
	root.Left.Right.Right = TNNew(4)

	root.Right = TNNew(1)
	root.Right.Left = TNNew(0)
	root.Right.Right = TNNew(8)

	fmt.Println(lowestCommonAncestor(root, p, q).Val)
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || p.Val == root.Val || q.Val == root.Val {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	return root
}
