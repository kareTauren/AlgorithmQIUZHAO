package home_work

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TNConstruct() *TreeNode {
	return &TreeNode{}
}

func TNNew(val int) *TreeNode {
	return &TreeNode{Val: val}
}
