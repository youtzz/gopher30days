## １４８·排序链表

题目链接：https://leetcode-cn.com/problems/sort-list/

### 描述

给你链表的头结点 `head` ，请将其按 **升序** 排列并返回 **排序后的链表** 。

### 样例

#### 样例１

```markdown
输入：head = [4,2,1,3]
输出：[1,2,3,4]
```

#### 样例２

```markdown
输入：head = [-1,5,3,4,0]
输出：[-1,0,3,4,5]
```

#### 样例３

```markdown
输入：head = []
输出：[]
```

### 提示

- 链表中节点的数目在范围 `[0, 5 * 104]` 内
- `-105 <= Node.val <= 105`

### 进阶

你可以在 `O(n log n)` 时间复杂度和常数级空间复杂度下，对链表进行排序吗？

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