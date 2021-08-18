## 题目


## 解题心得

二叉树的中序遍历，基础题。二叉树的遍历通常有两种写法，递归和迭代。

树的所有递归遍历其实就是一个模板：

```go
func order(root *TreeNode) {
    if root == nil {
       	return
    }
    // 先序遍历代码
    order(root.left)
    // 中序遍历代码
    order(root.right)
    // 后序遍历代码
}
```

迭代写法都要使用栈来存储节点。
