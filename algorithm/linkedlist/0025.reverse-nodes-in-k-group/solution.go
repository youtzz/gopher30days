package problem0025

import "github.com/sevenger/gopher30days/algorithm"

type ListNode = algorithm.ListNode

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}
	dummy := &ListNode{}
	pre := dummy
	for head != nil {
		start := head
		var i int
		for i = 1; i < k && head.Next != nil; i++ {
			head = head.Next
		}
		// 如果i!=k，直接连接剩下的节点并break
		if i != k {
			pre.Next = start
			break
		}

		// 断链
		next := head.Next
		head.Next = nil

		// 翻转
		reverse(start)

		// 连接
		pre.Next = head
		pre = start

		head = next
	}

	return dummy.Next
}

func reverse(head *ListNode) {
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
}
