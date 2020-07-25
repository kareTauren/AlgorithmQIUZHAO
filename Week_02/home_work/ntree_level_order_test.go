package home_work

import (
	"fmt"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	node := Construct()
	node.Val = 1

	node2 := New(3)
	node2.Children = []*Node{New(5), New(6)}

	node.Children = []*Node{node2, New(2), New(4)}
	fmt.Println(levelOrder(node))
}

func levelOrder(root *Node) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	queue := Queue{root}
	for queue.size() > 0 {
		size := queue.size()
		ansTemp := make([]int, 0)
		for i := 0; i < size; i++ {
			node := queue.remove()
			ansTemp = append(ansTemp, node.Val)
			queue.addAll(node.Children)
		}
		ans = append(ans, ansTemp)
	}
	return ans
}
