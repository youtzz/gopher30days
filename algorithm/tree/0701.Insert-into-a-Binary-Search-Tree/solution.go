package problem0701

import "github.com/sevenger/gopher30days/algorithm"

type TreeNode = algorithm.TreeNode

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	insertNode := &TreeNode{Val: val}
	if root == nil {
		return insertNode
	}

	node := root
	var prev *TreeNode // 保存上一个遍历过的一个树节点
	// 找到该插入的位置
	for node != nil {
		prev = node
		if val > node.Val {
			node = node.Right
		} else {
			node = node.Left
		}
	}
	if val > prev.Val {
		prev.Right = insertNode
	} else {
		prev.Left = insertNode
	}
	return root
}
