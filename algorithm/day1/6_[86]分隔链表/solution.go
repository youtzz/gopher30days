package main

import . "github.com/sevenger/gopher30days/algorithm"

// first solution
func partition(head *ListNode, x int) *ListNode {
	small, large := &ListNode{}, &ListNode{}
	smallNode, largeNode := small, large
	for head != nil {
		if head.Val < x {
			smallNode.Next = head
			smallNode = smallNode.Next
		} else {
			largeNode.Next = head
			largeNode = largeNode.Next
		}
		head = head.Next
	}
	largeNode.Next = nil
	smallNode.Next = large.Next
	return small.Next
}
