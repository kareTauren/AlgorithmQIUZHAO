package home_work

type Node struct {
	Val      int
	Children []*Node
}

func Construct() *Node {
	return &Node{}
}

func New(val int) *Node {
	return &Node{Val: val}
}
