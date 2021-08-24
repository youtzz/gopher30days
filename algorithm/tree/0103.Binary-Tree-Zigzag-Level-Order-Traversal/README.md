## １０３·二叉树的锯齿形层序遍历

题目链接：https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/

### 描述

给定一个二叉树，返回其节点值的锯齿形层序遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

例如：

给定二叉树 `[3,9,20,null,null,15,7]`,

```markdown
    3
   / \
  9  20
    /  \
   15   7
```

返回锯齿形层序遍历如下：

```markdown
[
  [3],
  [20,9],
  [15,7]
]
```

## 解题心得

也没什么好说的，层次遍历即可解题，用一个变量保存循环次数，通过这个变量决定每次是正向插入数据还是反向插入数据。
