package pratice

import (
	"fmt"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Construct() *TreeNode {
	treeNode := TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	}
	return &treeNode
}

func TestInorderTraversal(t *testing.T) {
	treeNode := Construct()
	treeNode.Val = 41
	treeNode.Right = &TreeNode{
		Val: 65,
		Left: &TreeNode{
			Val:   50,
			Left:  nil,
			Right: nil,
		},
		Right: nil,
	}
	treeNode.Left = &TreeNode{
		Val:   20,
		Left:  &TreeNode{Val: 11},
		Right: &TreeNode{Val: 29},
	}
	// input 1，null，2，3
	// inOrder reverse result is 1, 3, 2
	fmt.Println(inorderTraversal(treeNode))
}

// 中序遍历 左 - 根 - 右
func inorderTraversal(root *TreeNode) []int {
	rs := []int{}
	helper(root, &rs)
	return rs
}

func helper(root *TreeNode, rs *[]int) {
	if root == nil {
		return
	}
	helper(root.Left, rs)
	*rs = append(*rs, root.Val)
	helper(root.Right, rs)
}

type stack []*TreeNode

func (s *stack) Push(node *TreeNode) {
	*s = append(*s, node)
}

func (s *stack) Pop() *TreeNode {
	n := []*TreeNode(*s)[len(*s)-1]
	*s = []*TreeNode(*s)[:len(*s)-1]
	return n
}

func TestInOrderReverse(t *testing.T) {
	treeNode := Construct()
	treeNode.Val = 2
	treeNode.Right = &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
		Right: nil,
	}
	treeNode.Left = &TreeNode{Val: 1}
	fmt.Println(inOrderReverse2(treeNode))
}

func inOrderReverse2(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}
	cur := root
	sk := stack{}
	for cur != nil || len(sk) > 0 {
		for cur != nil {
			sk.Push(cur)
			cur = cur.Left
		}
		cur = sk.Pop()
		ans = append(ans, cur.Val)
		cur = cur.Right
	}
	return ans
}
