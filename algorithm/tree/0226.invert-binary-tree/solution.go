package problem0226

import "github.com/sevenger/gopher30days/algorithm"

type TreeNode = algorithm.TreeNode

// 递归法，本质上是个先序遍历
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left

	invertTree(root.Left)
	invertTree(root.Right)

	return root
}

// 广度优先，注意nil节点也要保存下来
func invertTree_bfs(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var res [][]*TreeNode
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		var list []*TreeNode
		l := len(queue)
		for i := 0; i < l; i++ {
			node := queue[0]
			queue = queue[1:]
			list = append(list, node)
			if node != nil {
				queue = append(queue, node.Left, node.Right)
			}
		}
		res = append(res, list)
	}

	for i := 0; i < len(res)-1; i++ {
		deal, use := res[i], res[i+1]
		for len(deal) > 0 {
			node := deal[len(deal)-1]
			deal = deal[:len(deal)-1]
			if node == nil {
				continue
			}
			node.Left, use = use[len(use)-1], use[:len(use)-1]
			node.Right, use = use[len(use)-1], use[:len(use)-1]
		}
	}

	return root
}
