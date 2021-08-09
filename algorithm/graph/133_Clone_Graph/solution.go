package clonegraph

import . "github.com/sevenger/gopher30days/algorithm"

// first solution
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}
	if head.Next == nil {
		return true
	}

	// 使用快慢指针将前一半回文存入栈
	stack := make([]*ListNode, 0)
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		stack = append(stack, slow)
		fast = fast.Next.Next
		slow = slow.Next
	}

	// fast不为nil说明链表是奇数个，slow指针在中间位置，要往前进一位
	if fast != nil {
		slow = slow.Next
	}

	// 不断退栈比较回文
	for slow != nil {
		val := slow.Val
		if stack[len(stack)-1].Val == val {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
		slow = slow.Next
	}

	return true
}
