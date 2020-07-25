package home_work

import (
	"fmt"
	"testing"
)

func TestPostOrder(t *testing.T) {
	node := Construct()
	node.Val = 1

	node2 := New(3)
	node2.Children = []*Node{New(5), New(6)}

	node.Children = []*Node{node2, New(2), New(4)}

	fmt.Println(postOrderTraverse(node))
}

func postOrderTraverse(node *Node) []int {
	ans := make([]int, 0)
	if node == nil {
		return ans
	}
	stack := Stack{node}
	for stack.Size() > 0 {
		node := stack.Pop()
		ans = append([]int{node.Val}, ans...)
		for i := 0; i < len(node.Children); i++ {
			stack.Push(node.Children[i])
		}
	}
	return ans
}
