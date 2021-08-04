## 题目



## 解题心得

这题是上一题的变种，也是链表逆置，区别是要求[left，right]范围内的链表节点逆置。

我第一时间想到的就是先遍历，将left-right范围内的链表截取出来，然后对截取的片段进行逆置，最后再拼接回去。

代码很轻松的就完成了，但是缺点是用了截取链表时用了两次循环。查看题解后，发现其实是可以用一次遍历就完成这道题目，要使用一种叫穿针引线的方法。

一共用到三个指针，pre、cur、next。pre是个dummyNode，永远指向left位置的节点的前一位，cur指向已逆置的节点的最后一个，next指向需要逆置的节点。

1. next := cur.next
2. cur.next = next.next
3. next.next = pre.next
4. pre.next = next

由于next指针成为了心得pre.next，而cur始终指向已逆置的节点的最后一个，所以当下一轮循环时，next:=cur.next可以使得next是需要逆置的节点。