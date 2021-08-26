package problem0416

import "math/bits"

// 第一次解：循环32次，每次取x,y的最后一位，统计不相同的数
func hammingDistance(x int, y int) int {
	count := 0
	for i := 0; i < 32; i++ {
		bitx := x >> i & 1
		bity := y >> i & 1
		if bitx != bity {
			count++
		}
	}
	return count
}

// 利用golang内置函数统计1的数量
func hammingDistance2(x int, y int) int {
	return bits.OnesCount(uint(x ^ y))
}

// 利用汉明距离公式 z&(z-1)得到z去除最后一个1的值
func hammingDistance_best(x int, y int) int {
	count := 0
	z := x ^ y
	for z != 0 {
		z = z & (z - 1)
		count++
	}
	return count
}
