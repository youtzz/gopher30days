## ３５·搜索插入位置

题目链接：https://leetcode-cn.com/problems/search-insert-position/

### 描述

给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

请必须使用时间复杂度为 O(log n) 的算法。

### 样例

#### 样例１

```markdown
输入: nums = [1,3,5,6], target = 5
输出: 2
```
#### 样例２

```markdown
输入: nums = [1,3,5,6], target = 2
输出: 1
```

#### 样例３

```markdown
输入: nums = [1,3,5,6], target = 2
输出: 1
```

#### 样例４

```markdown
输入: nums = [1,3,5,6], target = 0
输出: 0
```

#### 样例５

```markdown
输入: nums = [1], target = 0
输出: 0
```

### 提示

- `1 <= nums.length <= 104`
- `-104 <= nums[i] <= 104`
- `nums` 为**无重复元素的升序**排列数组
- `-104 <= target <= 104`

## 解题心得

这题也是用二分法，但是题目要求没找到时返回应该插入的位置，所以用二分法的[第二种通用模板](https://leetcode-cn.com/leetbook/read/binary-search/xerqxt/)就可以了。

由于要求返回应该插入的位置，所以用golang库里的的sort.Search函数也是可以的。

对于二分法的模板，可以详细阅读[这篇文章](https://blog.csdn.net/xiao_jj_jj/article/details/106018702)。
