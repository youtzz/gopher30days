package problem0160

import "github.com/sevenger/gopher30days/algorithm"

type ListNode = algorithm.ListNode

// 最容易想到的解法，用hash存储下第一条链表所有的节点，遍历第二条链表时检查节点在hash中是否已存在
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	hash := make(map[*ListNode]bool)
	for headA != nil {
		hash[headA] = true
		headA = headA.Next
	}
	for headB != nil {
		if hash[headB] {
			return headB
		}
		headB = headB.Next
	}
	return nil
}

// 我的解法2：没有使用多余空间，通过先算出list1和list2的长度的差值，使长的那条链到达起始位置，最后用双指针遍历，找到l==r的节点。
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	// 分别计算长度
	var len1, len2 int
	list1, list2 := headA, headB
	for list1 != nil {
		list1 = list1.Next
		len1++
	}
	for list2 != nil {
		list2 = list2.Next
		len2++
	}

	// 找到较长的那条链
	minus := len1 - len2
	if minus < 0 {
		minus *= -1
		headA, headB = headB, headA
	}
	// 让较长的链和另一条一样长
	list1 = headA
	for i := 0; i < minus; i++ {
		list1 = list1.Next
	}

	// 寻找两条链的第一个交点
	list2 = headB
	for i := 0; i < len1; i++ {
		if list1 == list2 {
			break
		}
		list1 = list1.Next
		list2 = list2.Next
	}
	if list1 != list2 {
		return nil
	}
	return list1
}

// 最佳解法，使用双指针
func getIntersectionNode_best(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pa, pb := headA, headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
}
