package main

import . "github.com/sevenger/gopher30days/algorithm"

// 递归法
func inorderTraversal(root *TreeNode) []int {
	rst := make([]int, 0)
	inorder(root, &rst)
	return rst
}

func inorder(root *TreeNode, rst *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, rst)
	*rst = append(*rst, root.Val)
	inorder(root.Right, rst)
}

// 遍历法
func inorderTraversal_no_recusive(root *TreeNode) (rst []int) {
	var stack []*TreeNode
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		rst = append(rst, root.Val)
		root = root.Right
	}
	return
}
