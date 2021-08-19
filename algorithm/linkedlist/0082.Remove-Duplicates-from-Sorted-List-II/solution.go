package problem0082

import "github.com/sevenger/gopher30days/algorithm"

type ListNode = algorithm.ListNode

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{Next: head}
	prev, cur := dummy, head

	for cur != nil {
		// 遍历到最后一个值和cur相同的节点
		for cur.Next != nil && cur.Val == cur.Next.Val {
			cur = cur.Next
		}
		// 如果上一步cur没有动，prev向下移动
		if prev.Next == cur {
			prev = prev.Next
		} else {
			// 上一步cur变了，prev.Next变动
			prev.Next = cur.Next
		}

		cur = cur.Next
	}
	return dummy.Next
}
