package numberof1bits

// first solution
func hammingWeight(num uint32) int {
	var count uint32
	for i := 0; i < 32; i++ {
		count += num >> i & 1
	}
	return int(count)
}

// best solution
func hammingWeight_best(num uint32) int {
	var count int
	for num != 0 {
		count++
		num = num & (num - 1)
	}
	return count
}
