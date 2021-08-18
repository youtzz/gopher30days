package problem94

import "github.com/sevenger/gopher30days/algorithm"

type TreeNode = algorithm.TreeNode

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
	print(1, 2, 3, 4, func(i int) {
		i += 1000
	})
	print(1, 2, 3, 4, func(i int) {
		i -= 10000
	})
	return res
}

func print(a, b, c, d int, f func(i int)) {
	if a > b {
		f(a)
	}
}
