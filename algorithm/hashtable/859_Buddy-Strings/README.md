## ８５９·亲密字符串

题目链接：https://leetcode-cn.com/problems/buddy-strings/

### 描述

给定两个由小写字母构成的字符串 `A` 和 `B` ，只要我们可以通过交换 `A` 中的两个字母得到与 `B` 相等的结果，就返回 `true` ；否则返回 `false` 。

交换字母的定义是取两个下标 `i` 和 `j` （下标从 `0` 开始），只要 i!=j 就交换 `A[i]` 和 `A[j]` 处的字符。例如，在 `"abcd"` 中交换下标 `0` 和下标 `2` 的元素可以生成 `"cbad"` 。

### 样例

### 样例１

```markdown
输入： A = "ab", B = "ba"
输出： true
解释： 你可以交换 A[0] = 'a' 和 A[1] = 'b' 生成 "ba"，此时 A 和 B 相等。
```
### 样例２

```markdown
输入： A = "ab", B = "ab"
输出： false
解释： 你只能交换 A[0] = 'a' 和 A[1] = 'b' 生成 "ba"，此时 A 和 B 不相等。
```

### 样例３

```markdown
输入： A = "aa", B = "aa"
输出： true
解释： 你可以交换 A[0] = 'a' 和 A[1] = 'a' 生成 "aa"，此时 A 和 B 相等。
```

### 样例４

```markdown
输入： A = "aaaaaaabc", B = "aaaaaaacb"
输出： true
```

### 样例５

```markdown
输入： A = "", B = "aa"
输出： false
```

## 解题心得

这题挺有意思的，看着挺复杂，其实核心内容拆解后很简单。就是求字符串两个不相同的index的字符串，通过交换位置，能不能组合成相同的字符串。因为题目有要求，(aa, aa)这两个字符串返回true，（aab,aab）这两个字符串返回false，前者是因为可以交换，后者是因为不能交换。

我使用了hashset统计一个字符串的字典，使用`len(s1)>len(hashset)`可以判断（aa，aa）这种情况。