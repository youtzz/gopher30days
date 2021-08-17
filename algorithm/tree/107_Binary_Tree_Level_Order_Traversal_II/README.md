## １０７·二叉树的层序遍历 II

题目链接：https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii/

### 描述

给定一个二叉树，返回其节点值自底向上的层序遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）

例如：
给定二叉树 `[3,9,20,null,null,15,7]`,

```markdown
    3
   / \
  9  20
    /  \
   15   7
```

返回其自底向上的层序遍历为：

```markdown
[
  [15,7],
  [9,20],
  [3]
]
```

## 解题心得

和107题没有什么区别，遍历完成后把切片反转就可以了。

```go
l := len(res)
for i := 0; i < l/2; i++ {
    res[i], res[l-1-i] = res[l-1-i], res[i]
}
```