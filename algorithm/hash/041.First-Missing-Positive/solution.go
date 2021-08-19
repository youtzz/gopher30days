package problem0041

func firstMissingPositive(nums []int) int {
	set := make(map[int]bool, len(nums))
	for _, v := range nums {
		if v >= 0 {
			set[v] = true
		}
	}

	var i int
	for i = 1; set[i]; i++ {
	}
	return i
}
