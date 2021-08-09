## ６1·搜索区间

题目链接：https://www.lintcode.com/problem/61/description

### 描述

给定一个包含 n 个整数的排序数组，找出给定目标值 target 的起始和结束位置。

如果目标值不在数组中，则返回[-1, -1]

### 样例

### 样例1

输入：

```markdown
数组 = []
target = 9
```

输出：

```markdown
[-1,-1]
```

解释：

9不在数组中。

### 样例2

输入：

```markdown
数组 = [5, 7, 7, 8, 8, 10]
target = 8
```

输出：

```markdown
[3，4]
```

解释：

数组的[3,4]子区间值为8。

### 挑战

时间复杂度 O(log *n*)


## 解题心得

这题比较有意思，刚好可以用二分算法的第三种模板来解决。第三种模板的特点就是可以在遍历中通过控制`nums[mid] == target`时的操作来决定最终找到target的是第一次出现的还是最后一次出现的。

```go
for left+1 < right {
    ...
    if nums[mid] == target {
        right = mid // 查找第一次出现的target下标
    }
    ...
}

for left+1 < right {
    ...
    if nums[mid] == target {
        left = mid　// 查找最后一次出现的target下标
    }
    ...
}
```

