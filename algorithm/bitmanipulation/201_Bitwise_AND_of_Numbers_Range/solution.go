package bitwiseandofnumbersrange

// best solution
func rangeBitwiseAnd(left int, right int) int {
	shift := 0
	for left != right {
		left, right = left>>1, right>>1
		shift++
	}
	return right << shift
}

// best solution2
func rangeBitwiseAnd2(left int, right int) int {
	for left < right {
		// right不断失去最靠后的1，所以当right<left时，此时的right就是公共前缀
		right &= right - 1
	}
	return right
}
