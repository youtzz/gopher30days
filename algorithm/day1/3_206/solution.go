package main

import . "github.com/sevenger/gopher30days/algorithm"

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dummy := &ListNode{}
	for head != nil {
		node := head
		head = head.Next
		node.Next = dummy.Next
		dummy.Next = node
	}
	return dummy.Next
}