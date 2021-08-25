package problem0101

import "math"

import "github.com/sevenger/gopher30days/algorithm"

type TreeNode = algorithm.TreeNode

// 第一次解，使用广度优先遍历树，分析每一次遍历的节点的值是否对称
func isSymmetric(root *TreeNode) bool {
	const null = math.MinInt32
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		var list []int
		l := len(queue)
		for i := 0; i < l; i++ {
			node := queue[0]
			queue = queue[1:]
			// nil的时候存null值
			if node == nil {
				list = append(list, null)
				continue
			} else if node == root {
				// 头结点多存一个值，防止后面的逻辑误判断
				list = append(list, node.Val, node.Val)
			} else {
				list = append(list, node.Val)
			}
			queue = append(queue, node.Left, node.Right)
		}

		// 如果list的长度是奇数，一定不是对称的
		if len(list)%2 != 0 {
			return false
		}
		// 不断比对左右两个数
		for l, r := 0, len(list); l != r; l, r = l+1, r-1 {
			if list[l] != list[r-1] {
				return false
			}
		}
	}
	return true
}

// 最佳解法，迭代法
func isSymmetric_best(root *TreeNode) bool {
	queue := []*TreeNode{root, root}
	for len(queue) > 0 {
		q, v := queue[0], queue[1]
		queue = queue[2:]
		if q == nil && v == nil {
			continue
		}
		if q == nil || v == nil {
			return false
		}
		if q.Val != v.Val {
			return false
		}
		queue = append(queue, q.Left, v.Right, q.Right, v.Left)
	}
	return true
}

// 最佳解法，递归法
func isSymmetric_best_recur(root *TreeNode) bool {
	return check(root, root)
}

func check(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && check(p.Left, q.Right) && check(p.Right, q.Left)
}
