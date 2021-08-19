package problem0206

import "github.com/sevenger/gopher30days/algorithm"

type ListNode = algorithm.ListNode

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

func reverseList_recur(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := reverseList_recur(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
