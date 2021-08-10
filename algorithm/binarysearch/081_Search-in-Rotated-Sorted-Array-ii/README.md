## ８１·搜索旋转排序数组 II

题目链接：https://leetcode-cn.com/problems/search-in-rotated-sorted-array-ii/

### 描述

已知存在一个按非降序排列的整数数组 `nums` ，数组中的值不必互不相同。

在传递给函数之前，`nums` 在预先未知的某个下标 `k`（`0 <= k < nums.length`）上进行了 **旋转** ，使数组变为 `[nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]`（下标 **从 0 开始** 计数）。例如， `[0,1,2,4,4,4,5,6,6,7]` 在下标 5 处经旋转后可能变为 [4,5,6,6,7,0,1,2,4,4] 。

给你 **旋转后** 的数组 `nums` 和一个整数 `target` ，请你编写一个函数来判断给定的目标值是否存在于数组中。如果 `nums` 中存在这个目标值 `target` ，则返回 `true` ，否则返回 `false` 。

### 样例

### 样例１

```markdown
输入：nums = [2,5,6,0,0,1,2], target = 0
输出：true
```
### 样例２

```markdown
输入：nums = [2,5,6,0,0,1,2], target = 3
输出：false
```

### 进阶

- 这是 [搜索旋转排序数组](https://leetcode-cn.com/problems/search-in-rotated-sorted-array/description/) 的延伸题目，本题中的 `nums` 可能包含重复元素。
- 这会影响到程序的时间复杂度吗？会有怎样的影响，为什么？

## 解题心得

这题和33题是一样的套路，只是这次数组里有重复的数字。例如nums=[1，0，1]，nums[left]=nums[right]=nums[mid]=1，此时无法区分哪边数组是有序的。解决的方法也不难，相等时，让left++，right--后再进行下一轮比较即可。