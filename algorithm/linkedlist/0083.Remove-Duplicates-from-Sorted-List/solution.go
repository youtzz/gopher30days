package problem0083

import . "github.com/sevenger/gopher30days/algorithm"

// 我的解法1：用hashSet存下每次遍历的数据，如果hashSet已有说明数据重复
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	set := map[int]bool{}
	dummy := &ListNode{Next: head}
	prev, cur := dummy, dummy.Next

	for cur != nil {
		if set[cur.Val] {
			prev.Next = cur.Next
		} else {
			set[cur.Val] = true
			prev = prev.Next
		}
		cur = cur.Next
	}
	return dummy.Next
}

// 我的解法2：遍历链表，每次都将指针遍历到重复的最后一个节点，再重新链接链表
func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{Next: head}
	prev, cur := dummy, dummy.Next

	for cur != nil {
		for cur.Next != nil && cur.Val == cur.Next.Val {
			cur = cur.Next
		}

		// 如果cur发生了移动（说明有重复），重新链接prev的next指针
		if prev.Next != cur {
			prev.Next = cur
		}
		prev = prev.Next
		cur = cur.Next
	}
	return dummy.Next
}

// 递归解法
func deleteDuplicates_recur(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	head = deleteDuplicates(head.Next)

	if head.Val == head.Next.Val {
		head = head.Next
	}
	return head
}

// 最佳解法
func deleteDuplicates_best(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}
