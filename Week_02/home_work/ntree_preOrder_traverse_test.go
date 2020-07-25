package home_work

import (
	"fmt"
	"testing"
)

func TestPreOrderTraverse(t *testing.T) {
	node := Construct()
	node.Val = 1

	node2 := New(3)
	node2.Children = []*Node{New(5), New(6)}

	node.Children = []*Node{node2, New(2), New(4)}

	fmt.Println(preOrderTraverse(node))
}

func preOrderTraverse(node *Node) []int {
	ans := make([]int, 0)
	if node == nil {
		return ans
	}

	stack := Stack{node}
	for stack.Size() > 0 {
		node := stack.Pop()
		ans = append(ans, node.Val)
		for i := len(node.Children) - 1; i >= 0; i-- {
			stack.Push(node.Children[i])
		}
	}
	return ans
}
