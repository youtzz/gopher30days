package main

// first solution
type MinStack struct {
	head *Node
	min  *Node
}

type Node struct {
	Val  int
	Next *Node
}

func Constructor() MinStack {
	return MinStack{
		head: &Node{},
		min:  nil,
	}
}

func (this *MinStack) Push(val int) {
	node := &Node{Val: val}
	node.Next = this.head.Next
	this.head.Next = node

	if this.min == nil || val < this.min.Val {
		this.min = node
	}
}

func (this *MinStack) Pop() {
	node := this.head.Next
	this.head.Next = node.Next
	// 重新寻找最小节点
	if this.min == node {
		cur := this.head.Next
		min := cur
		for cur != nil {
			if min.Val > cur.Val {
				min = cur
			}
			cur = cur.Next
		}
		this.min = min
	}
}

func (this *MinStack) Top() int {
	return this.head.Next.Val
}

func (this *MinStack) GetMin() int {
	return this.min.Val
}
