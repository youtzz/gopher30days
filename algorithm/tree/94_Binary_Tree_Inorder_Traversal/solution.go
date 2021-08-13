package binaryTreeInorderTraversal

import . "github.com/sevenger/gopher30days/algorithm"

func inorderTraversal(root *TreeNode) (res []int) {
	inorder(&res, root)
	return
}

func inorder(res *[]int, root *TreeNode) {
	if root == nil {
		return
	}
	inorder(res, root.Left)
	*res = append(*res, root.Val)
	inorder(res, root.Right)
}

func inorderTraversal_interactive(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	var stack []*TreeNode
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root, stack = stack[len(stack)-1], stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return res
}
