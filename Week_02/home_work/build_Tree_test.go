package home_work

import (
	"fmt"
	"testing"
)

func TestBuildTree(t *testing.T) {
	preOrder := []int{1, 2, 4, 5, 3, 6, 7}
	inOrder := []int{4, 2, 5, 1, 6, 3, 7}
	fmt.Println(buildTree(preOrder, inOrder))
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	// 使用前序数组的第一个值确定根节点(root)
	root := &TreeNode{Val: preorder[0]}
	// 找到中序数组的中间值，中序数组的中间值左边就是左子树，右边的值就是右子树的值
	i := 0
	for ; i < len(preorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}
