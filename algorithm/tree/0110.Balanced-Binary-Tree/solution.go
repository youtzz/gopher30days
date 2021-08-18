package problem0110

import "github.com/sevenger/gopher30days/algorithm"

type TreeNode = algorithm.TreeNode

// 我的解法
func isBalanced(root *TreeNode) bool {
	is, _ := dfs(root)
	return is
}

func dfs(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}
	isl, l := dfs(root.Left)
	isr, r := dfs(root.Right)
	// 如果不是平衡二叉树返回false
	if !isl || !isr {
		return false, 0
	}
	// 通过左子树深度-右子树深度判断是不是平衡二叉树
	reduce := l - r
	if reduce > 1 || reduce < -1 {
		return false, 0
	}
	return true, max(l, r) + 1
}

// 最佳解法
func isBalanced_best(root *TreeNode) bool {
	// 只要有一棵子树不是平衡二叉树，递归函数总是返回 -1
	return dfs_best(root) >= 0
}

func dfs_best(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := dfs_best(root.Left)
	r := dfs_best(root.Right)
	// 差值绝对值 > 1 return -1
	if l-r > 1 || l-r < -1 || l < 0 || r < 0 {
		return -1
	}
	return max(l, r) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
