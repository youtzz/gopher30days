## 题目



## 解题心得

这题是排序链表，题目要求nlogn的时间复杂度，我感觉有点难度，印象中要让时间复杂度到达nlogn的算法有快速排序和归并排序，链表搞快速排序有点不科学，归并排序我也不太熟悉。

最后尝试着写了，还是没有写出来，主要是归并中的归想不出该怎么用在链表上。并倒是很简单，和前面的力扣21题其实是一样的。

对链表使用归的难点主要是如何找到中间节点，并且将链表断开。后来我了解到该使用快慢指针。一共用到三个变量，brk、fast、slow

```go
var brk *ListNode
slow, fast := head, head
for fast != nil && fast.Next.Next != nil {
	fast = fast.Next.Next
	// 当fast即将到达
	if fast == nil || fast.Next == nil {
	    brk = slow // 保存slow之前的位置	
    }   
	slow = slow.Next
}
brk.Next = nil // 断链
```