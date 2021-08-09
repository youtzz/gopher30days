## 题目

## 解题心得

这道题本质上是二分查找，因为当前版本n如果是错误的，说明大于n的版本都是错误的，反之如果n是正确的，那么小于n的版本都是正确的。

所以只要设置left、right，取mid值判断，然后根据mid来设置新的left或right就能最快找出答案。

这题的二分算法用golang内置的函数库也可以直接解决。