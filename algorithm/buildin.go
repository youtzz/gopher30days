package algorithm

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewLinkList(args ...int) *ListNode {
	head := ListNode{}
	for _, v := range args {
		node := ListNode{Val: v}
		head.Next = &node
		head = *head.Next
	}
	return head.Next
}

func CompareLinkList(l1, l2 *ListNode) (rst bool) {
	switch {
	case l1 == nil && l2 == nil:
		return true
	case l1 == nil && l2 != nil:
		fallthrough
	case l1 != nil && l2 == nil:
		return false
	default:
		for l1 != nil && l2 != nil {
			if l1.Val != l2.Val {
				return rst
			}
			l1 = l1.Next
			l2 = l2.Next
		}
	}

	if l1 == nil && l2 == nil {
		return true
	}

	for l1 != nil || l2 != nil {
	}

	return
}
