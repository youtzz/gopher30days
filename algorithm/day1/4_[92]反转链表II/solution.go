package main

import . "github.com/sevenger/gopher30days/algorithm"

// first solution
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || left > right || left == right {
		return head
	}

	dummyNode := &ListNode{Next: head}
	pre := dummyNode
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}

	rightNode := pre
	for i := 0; i < right-left+1; i++ {
		rightNode = rightNode.Next
	}

	leftNode, cur := pre.Next, rightNode.Next
	pre.Next, rightNode.Next = nil, nil
	reverseLinkList(leftNode)

	pre.Next, leftNode.Next = rightNode, cur
	return dummyNode.Next
}

func reverseLinkList(head *ListNode) {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
}

// best solution
func reverseBetweenBest(head *ListNode, left int, right int) *ListNode {
	if head == nil || left > right || left == right {
		return head
	}

	dummyNode := &ListNode{Next: head}
	pre := dummyNode

	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	cur := pre.Next
	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	return dummyNode.Next
}
