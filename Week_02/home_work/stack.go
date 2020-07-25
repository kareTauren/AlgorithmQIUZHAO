package home_work

type Stack []*Node

func (s *Stack) Push(node *Node) {
	*s = append(*s, node)
}

func (s *Stack) Pop() *Node {
	node := []*Node(*s)[len(*s)-1]
	*s = []*Node(*s)[:len(*s)-1]
	return node
}

func (s *Stack) Size() int {
	return len(*s)
}
