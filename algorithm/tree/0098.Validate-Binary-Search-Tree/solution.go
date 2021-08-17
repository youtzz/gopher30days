package problem0098

import (
	"github.com/sevenger/gopher30days/algorithm"
	"math"
)

type TreeNode = algorithm.TreeNode

// 我的解法：利用一棵二叉搜索树的中序遍历结果一定是有序的特性
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return false
	}

	var res []int
	inorder(&res, root)
	pre := math.MinInt64
	for _, v := range res {
		if pre >= v {
			return false
		}
		pre = v
	}
	return true
}

func inorder(res *[]int, root *TreeNode) {
	if root == nil {
		return
	}
	inorder(res, root.Left)
	*res = append(*res, root.Val)
	inorder(res, root.Right)
}

// 我的解法2：根据每一棵二叉搜索树的子树都是二叉搜索树的特性，可以利用分治思想
func isValidBST_dac(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *TreeNode, low, high int) bool {
	if root == nil {
		return true
	}

	if root.Val <= low || root.Val >= high {
		return false
	}
	return helper(root.Left, low, root.Val) && helper(root.Right, root.Val, high)
}
