package main

import . "github.com/sevenger/gopher30days/algorithm"

// 思路，用一个map存储每次遍历的val，如果map中存在就删除
func deleteDuplicates(head *ListNode) *ListNode {
	hashTable := make(map[int]struct{}, 0)
	hashTable[head.Val] = struct{}{}
	node := head
	for node != nil && node.Next != nil {
		val := node.Next.Val
		if _, ok := hashTable[val]; ok {
			node.Next = node.Next.Next
		} else {
			hashTable[val] = struct{}{}
		}
		node = node.Next
	}
	return head
}
