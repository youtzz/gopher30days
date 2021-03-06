## １１０·平衡二叉树

题目链接：https://leetcode-cn.com/problems/balanced-binary-tree/

### 描述

给定一个二叉树，判断它是否是高度平衡的二叉树。

本题中，一棵高度平衡二叉树定义为：

> 一个二叉树*每个节点* 的左右两个子树的高度差的绝对值不超过 1 。

### 样例

#### 样例１

```markdown
输入：root = [3,9,20,null,null,15,7]
输出：true
```
#### 样例２

```markdown
输入：root = [1,2,2,3,3,null,null,4,4]
输出：false
```

#### 样例３

```markdown
输入：root = []
输出：true
```

### 提示

- 树中的节点数在范围 `[0, 5000]` 内
- `-104 <= Node.val <= 104`

## 解题心得

这题可以用自顶向上和自顶向下来做，有了 `104` 题的经验，我第一时间想到的就是自定向下法。因为平衡二叉树的每一个子树都是平衡二叉树，所以每个树节点的左右子树都要进行比较。

我的递归函数有两个返回值 `(isBalanced bool, depth int)` ，每次递归都会判断上一次的结果，如果是返回了false说明不是平衡二叉树，就不递归了，否则返回两棵子树深度最大的值。

由于我的递归函数有两个返回值，所以实现的不是很优雅，参考了下标准答案，发现每次递归返回子树深度就可以了。如果出现非平衡二叉树，那么之后的递归每次都返回 `-1` ，最终结果一定是 小于 `0` 的。如果是平衡二叉树，那么最后返回的就是树的深度。所以最终递归结束后判断一下返回的 `depth` 就可以了。

