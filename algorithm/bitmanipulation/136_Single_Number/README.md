## 题目



## 解题心得

本题求无序数组里未重复的数字，数字最多重复两次。我第一反应是用遍历数组的时候用HashTable存数字，如果数字不存在就存入HashTable，存在的话就删除HashTable中的key。遍历结束后HashTable中唯一剩下的就是未重复的数字。

但其是最简单的方法是遍历时对数字进行异或操作，因为a^b^b=a，相同的数字会在异或时所以遍历最后算出的数字就是不重复的数字。~~(太神奇了)~~

异或操作有如下特征：

1. 任何数和0异或，结果为原来的数：a ^ 0 = a
2. 任何数和它自身异或，结果为0：a ^ a = 0
3. 异或运算满足交换律和结合律：a^b^a = a^a^b = a^(b^a)