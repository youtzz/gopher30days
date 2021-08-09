package countingbits

// first solution
func countBits(n int) []int {
	rst := make([]int, 0, n)
	for i := 0; i <= n; i++ {
		num, count := i, 0
		for num > 0 {
			count++
			num &= num - 1
		}
		rst = append(rst, count)
	}
	return rst
}

// best solution
func countBits_best(n int) []int {
	res := make([]int, n+1)

	for i := 1; i <= n; i++ {
		res[i] = res[i&(i-1)] + 1
	}
	return res
}
