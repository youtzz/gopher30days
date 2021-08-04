package main

import . "github.com/sevenger/gopher30days/algorithm"

// 思路：使用两个指针，一个dummy指针指向头，另一个指针不断往下遍历，只要
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	preHead := &ListNode{Next: head}
	pre, cur := preHead, preHead.Next
	for cur != nil {
		for cur.Next != nil && cur.Val == cur.Next.Val {
			cur = cur.Next
		}
		if pre.Next == cur {
			pre = pre.Next
		} else {
			pre.Next = cur.Next
		}
		cur = cur.Next
	}
	return preHead.Next
}

