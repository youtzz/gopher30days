//存在一个按升序排列的链表，给你这个链表的头节点 head ，请你删除所有重复的元素，使每个元素 只出现一次 。
//
// 返回同样按升序排列的结果链表。
//
//
//
// 示例 1：
//
//
//输入：head = [1,1,2]
//输出：[1,2]
//
//
// 示例 2：
//
//
//输入：head = [1,1,2,3,3]
//输出：[1,2,3]
//
//
//
//
// 提示：
//
//
// 链表中节点数目在范围 [0, 300] 内
// -100 <= Node.val <= 100
// 题目数据保证链表已经按升序排列
//
// Related Topics 链表
// 👍 615 👎 0

package main

type ListNode struct {
Val int
Next *ListNode
}

func main() {
head := &ListNode{
Val:  1,
Next: &ListNode{
Val:  1,
Next: &ListNode{
Val:  2,
Next: &ListNode{
Val:  3,
Next: &ListNode{
Val:  3,
Next: nil,
},
},
},
},
}
deleteDuplicates(head)
}
// 思路，用一个map存储每次遍历的val，如果map中存在就删除
//leetcode submit region begin(Prohibit modification and deletion)
/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
  */
  func deleteDuplicates(head *ListNode) *ListNode {
  if head == nil {
  return head
  }
  hashTable := make(map[int]struct{}, 0)
  hashTable[head.Val] = struct{}{}
  node := head
  for node != nil && node.Next != nil {
  val := node.Next.Val
  if _, ok := hashTable[val]; ok {
  node.Next = node.Next.Next
  } else {
  hashTable[val] = struct{}{}
  node = node.Next
  }
  }
  return head
  }
  //leetcode submit region end(Prohibit modification and deletion)
