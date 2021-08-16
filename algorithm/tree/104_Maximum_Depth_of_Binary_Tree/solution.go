package Maximum_Depth_of_Binary_Tree

import . "github.com/sevenger/gopher30days/algorithm"

func maxDepth_bfs(root *TreeNode) int {
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

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := maxDepth(root.Left)
	r := maxDepth(root.Right)
	return max(l, r) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
