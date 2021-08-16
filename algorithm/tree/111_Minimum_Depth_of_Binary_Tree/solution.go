package Minimum_Depth_of_Binary_Tree

import (
	. "github.com/sevenger/gopher30days/algorithm"
	"math"
)

// 我的解法 DFS
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	var depth int
	for len(queue) > 0 {
		depth++
		l := len(queue)
		for i := 0; i < l; i++ {
			node := queue[0]
			queue = queue[1:]
			// left和right都为nil说明是叶子节点
			if node.Left == nil && node.Right == nil {
				return depth
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return depth
}

// 标准解法 BFS
func minDepth_DFS(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}

	minD := math.MaxInt64
	if root.Left != nil {
		minD = min(minDepth(root.Left), minD)
	}
	if root.Right != nil {
		minD = min(minDepth(root.Right), minD)
	}
	return minD + 1
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
