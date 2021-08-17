package Binary_Tree_Level_Order_Traversal_II

import . "github.com/sevenger/gopher30days/algorithm"

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var res [][]int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		var list []int
		l := len(queue)
		for i := 0; i < l; i++ {
			node := queue[0]
			queue = queue[1:]
			list = append(list, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, list)
	}
	len := len(res)
	for i := 0; i < len/2; i++ {
		res[i], res[len-1-i] = res[len-1-i], res[i]
	}
	return res
}
