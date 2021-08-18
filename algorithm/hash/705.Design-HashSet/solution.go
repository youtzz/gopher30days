package problem0705

// MyHashSet 使用链表实现哈希集合
type MyHashSet struct {
	set []LinkedList
}

const base = 769 // 散列数

func Constructor() MyHashSet {
	return MyHashSet{set: make([]LinkedList, base)}
}

func (s *MyHashSet) hash(key int) int {
	return key % base
}

func (s *MyHashSet) Add(key int) {
	if !s.Contains(key) {
		h := s.hash(key)
		s.set[h].Insert(key)
	}
}

func (s *MyHashSet) Remove(key int) {
	h := s.hash(key)
	for e := s.set[h].Front(); e != nil; e = e.Next {
		if e.Val == key {
			s.set[h].Remove(e)
		}
	}
}

func (s *MyHashSet) Contains(key int) bool {
	h := s.hash(key)
	for e := s.set[h].Front(); e != nil; e = e.Next {
		if e.Val == key {
			return true
		}
	}
	return false
}

// LinkedList 链表
type LinkedList struct {
	Root ListNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *LinkedList) Front() *ListNode {
	return l.Root.Next
}

func (l *LinkedList) Insert(val int) {
	node := &ListNode{Val: val, Next: l.Root.Next}
	l.Root.Next = node
}

func (l *LinkedList) Remove(node *ListNode) {
	if l.Root.Next == node {
		l.Root.Next = nil
		return
	}

	cur := l.Root.Next
	prv := &ListNode{Next: cur}
	for cur != nil {
		if cur == node {
			prv.Next = cur.Next
		}
		cur = cur.Next
		prv = prv.Next
	}
}
