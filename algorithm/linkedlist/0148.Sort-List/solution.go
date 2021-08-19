package sortList

import "github.com/sevenger/gopher30days/algorithm"

type ListNode = algorithm.ListNode

// best solution
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow, fast := head, head
	var brk *ListNode // 用于断链
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		if fast == nil || fast.Next == nil {
			brk = slow
		}
		slow = slow.Next
	}
	brk.Next = nil
	head1 := sortList(head)
	head2 := sortList(slow)

	return merge(head1, head2)
}

func merge(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 == nil {
		cur.Next = l2
	}
	if l2 == nil {
		cur.Next = l1
	}
	return dummy.Next
}
