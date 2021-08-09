## 题目



## 解题心得

这题和上一题差不多，都是删除链表中的元素，唯一不同的是本题不保留重复元素，要求重复就全部删除。

最开始的想法就是双指针，一个指向头，另一个指向尾，指向尾的不断往前寻找非重复节点。

写出来了，但是不够优雅，代码有点长，主要原因是头节点可能会被删除，所以我加了一些if-else判断这种情况。

后面看了题解，和我的思路大体相同，只是用了一个哑节点，即dummy节点来当作头节点，使用go语言的话直接 ```dummy := &ListNode{Next: head}```，最后函数return dummy.next，所以就算头节点删除了也不会产生影响，非常方便。

采用了这种方式后，代码精简了许多。
