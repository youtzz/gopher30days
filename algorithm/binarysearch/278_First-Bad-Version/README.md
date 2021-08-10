## 278·第一个错误的版本

题目链接：https://leetcode-cn.com/problems/first-bad-version/

### 描述

你是产品经理，目前正在带领一个团队开发新的产品。不幸的是，你的产品的最新版本没有通过质量检测。由于每个版本都是基于之前的版本开发的，所以错误的版本之后的所有版本都是错的。

假设你有 n 个版本 [1, 2, ..., n]，你想找出导致之后所有版本出错的第一个错误的版本。

你可以通过调用`bool isBadVersion(version)`接口来判断版本号 version 是否在单元测试中出错。实现一个函数来查找第一个错误的版本。你应该尽量减少对调用 API 的次数。

### 样例

### 样例１

```markdown
输入：n = 5, bad = 4
输出：4
解释：
调用 isBadVersion(3) -> false
调用 isBadVersion(5) -> true
调用 isBadVersion(4) -> true
所以，4 是第一个错误的版本。
```
### 样例２

```markdown
输入：n = 1, bad = 1
输出：1
```

## 解题心得

这道题本质上是二分查找，因为当前版本n如果是错误的，说明大于n的版本都是错误的，反之如果n是正确的，那么小于n的版本都是正确的。

所以只要设置left、right，取mid值判断，然后根据mid来设置新的left或right就能最快找出答案。

这题的二分算法用golang内置的函数库也可以直接解决。