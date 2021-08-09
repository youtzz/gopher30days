package algorithm

import (
	"fmt"
	"strconv"
	"strings"
)

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

func PrintLinkList(head *ListNode) {
	fmt.Println(getFormatLinkListString(head))
}

func getFormatLinkListString(head *ListNode) string {
	var stack []int
	for head != nil {
		stack = append(stack, head.Val)
		head = head.Next
	}

	var sb strings.Builder
	sb.Grow(len(stack))
	sb.WriteString("{ ")
	for _, v := range stack {
		sb.WriteString(strconv.Itoa(v))
		sb.WriteString(" -> ")
	}
	sb.WriteString("nil }")
	return sb.String()
}

// -----------------树

type Tree *TreeNode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
