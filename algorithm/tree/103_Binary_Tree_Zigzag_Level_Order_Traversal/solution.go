package Binary_Tree_Zigzag_Level_Order_Traversal

import . "github.com/sevenger/gopher30days/algorithm"

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var res [][]int
	queue := []*TreeNode{root}
	times := 1
	for len(queue) > 0 {
		var list []int
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
			if times%2 == 0 {
				list = append(list, 0)
				copy(list[1:], list)
				list[0] = node.Val
			} else {
				list = append(list, node.Val)
			}
		}
		res = append(res, list)
		times++
	}
	return res
}
