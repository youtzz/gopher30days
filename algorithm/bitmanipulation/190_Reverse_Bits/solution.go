package reversebits

// first solution
func reverseBits(num uint32) uint32 {
	var res uint32
	for i := 0; i < 32; i++ {
		n := num >> i & 1    // 获取最后一位
		res |= n << (31 - i) // 通过或操作赋值
	}
	return res
}

// first solution2
func reverseBits2(num uint32) uint32 {
	var total uint32
	pow := 31
	for num != 0 {
		total += num & 1 << pow
		pow--
		num >>= 1
	}
	return total
}
