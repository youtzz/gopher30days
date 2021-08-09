## 题目

## 解题心得

这题是上一题的变种，也是求没有重复的数字，但是重复数字出现的次数变成了三次。利用O(n)空间复杂度的方法很简单，用HashTable记录每次数字出现的次数，最后遍历HashTable，出现次数为1的数就是无重复的数字。

想要空间复杂度为O(1)，肯定是要用位运算了，但是本题重复数字出现了三次，不能直接用异或操作算结果，该怎么解，着实让我犯了难。

具体的做法是这样的：把每一个数字都当成32位(题目给的数字范围是int32)的二进制数字来看，把所有数字的二进制的其中一位全部加起来，一定是3N，或者3N+1，因为重复数字出现三次，如果是3N+1，说明这一位有未重复位，需要保留。把所有3N+1的位用 | 操作的结果就是最终答案。

进阶的做法基于上一种，但是用到了数字逻辑的知识，暂时没有搞明白，有空再来填坑。