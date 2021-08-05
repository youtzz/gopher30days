package algorithm

// -----------------链表

type LinkList *ListNode

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewLinkList(arg []int) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for _, v := range arg {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}

func CompareLinkList(l1, l2 *ListNode) bool {
	rst := true
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			rst = false
			break
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	if l1 != nil || l2 != nil {
		rst = false
	}

	return rst
}

// -----------------树

type Tree *TreeNode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

