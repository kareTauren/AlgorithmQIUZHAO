package home_work

type Queue []*Node

func (q *Queue) add(node *Node) bool {
	*q = append(*q, node)
	return true
}

func (q *Queue) remove() *Node {
	if len(*q) == 0 {
		return nil
	}

	first := []*Node(*q)[0]
	*q = []*Node(*q)[1:]
	return first
}

func (q *Queue) element() *Node {
	first := []*Node(*q)[0]
	return first
}

func (q *Queue) size() int {
	return len(*q)
}

func (q *Queue) addAll(nodes []*Node) bool {
	*q = append(*q, nodes...)
	return true
}
