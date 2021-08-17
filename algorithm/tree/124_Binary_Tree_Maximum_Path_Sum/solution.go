package Binary_Tree_Maximum_Path_Sum

import (
	. "github.com/sevenger/gopher30days/algorithm"
	"math"
)

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32
	var maxGain func(root *TreeNode) int
	maxGain = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftGain := max(0, maxGain(root.Left))
		rightGain := max(0, maxGain(root.Right))

		maxSum = max(maxSum, root.Val+leftGain+rightGain)
		return max(leftGain, rightGain) + root.Val
	}

	maxGain(root)
	return maxSum
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
