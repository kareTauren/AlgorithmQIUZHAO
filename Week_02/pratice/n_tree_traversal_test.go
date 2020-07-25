package pratice

import (
	"fmt"
	"testing"
)

type Node struct {
	Val      int
	Children []*Node
}

func TestNTreeTraversal(t *testing.T) {
	// node := NewNode(1)
	// node2 := NewNode(3)
	// node2.Children = []*Node{NewNode(5), NewNode(6)}
	// node.Children = []*Node{node2, NewNode(2), NewNode(4)}
	// // 后续
	// // fmt.Println("后序：", postorder(node))
	// // fmt.Println("中序：", inorder(node))
	// // fmt.Println("前序：", preOrder(node))
	//
	// fmt.Println(nTreeLayerTraversal(node))
}

func NewNode(val int) *Node {
	return &Node{Val: val}
}

// N 叉树遍历 后续遍历
func postorder(root *Node) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}

	stack := NodeStack{root}
	for len(stack) > 0 {
		node := stack.Pop()
		ans = append([]int{node.Val}, ans...)
		for i := 0; i < len(node.Children); i++ {
			stack = append(stack, node.Children[i])
		}
	}
	return ans
}

// N 叉树遍历 中序遍历 (未完成)
func inorder(root *Node) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	stack := NodeStack{root}
	for len(stack) > 0 {
		node := stack.Pop()
		ans = append(ans, node.Val)
		for i := 0; i < len(node.Children); i++ {
			stack.Push(node.Children[i])
		}
	}
	return ans
}

// 前序遍历
func preOrder(root *Node) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	stack := NodeStack{root}
	for len(stack) > 0 {
		node := stack.Pop()
		ans = append(ans, node.Val)
		for i := len(node.Children) - 1; i >= 0; i-- {
			stack.Push(node.Children[i])
		}
	}

	return ans
}

// 层序遍历
func nTreeLayerTraversal(root *Node) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}

	queue := NodeQueue{root}
	for queue.size() > 0 {
		tmpAns := []int{}
		size := queue.size()
		fmt.Println("size ", size)
		for i := 0; i < size; i++ {
			node := queue.remove()
			tmpAns = append(tmpAns, node.Val)
			queue.addAll(node.Children)
		}
		ans = append(ans, tmpAns)
	}
	return ans
}

type NodeStack []*Node

func (s *NodeStack) Push(node *Node) {
	*s = append(*s, node)
}

func (s *NodeStack) Pop() *Node {
	n := []*Node(*s)[len(*s)-1]
	*s = []*Node(*s)[:len(*s)-1]
	return n
}

type NodeQueue []*Node

func (q *NodeQueue) add(node *Node) bool {
	*q = append(*q, node)
	return true
}

func (q *NodeQueue) remove() *Node {
	if len(*q) == 0 {
		return nil
	}
	first := []*Node(*q)[0]
	*q = []*Node(*q)[1:]
	return first
}

func (q *NodeQueue) element() *Node {
	first := []*Node(*q)[0]
	return first
}

func (q *NodeQueue) size() int {
	return len(*q)
}

func (q *NodeQueue) addAll(nodes []*Node) bool {
	*q = append(*q, nodes...)
	return true
}
