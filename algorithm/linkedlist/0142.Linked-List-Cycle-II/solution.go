package problem0142

import "github.com/sevenger/gopher30days/algorithm"

type ListNode = algorithm.ListNode

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == head {
		return head
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		// 快指针走过的路程是慢指针的两倍
		// 所以当快慢指针相遇时，设置一个从头结点开始的指针p
		// p 和 slow 将在环入口相遇
		if fast == slow {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}
