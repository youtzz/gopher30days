package problem0080

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	index := 2
	for i := 2; i < len(nums); i++ {
		if nums[index-2] != nums[i] {
			nums[index] = nums[i]
			index++
		}
	}
	return index
}

// 数组去重保留K个通用模板
func removeDuplicates_best(nums []int) int {
	var process func(k int) int
	process = func(k int) int {
		u := 0
		for _, v := range nums {
			if u < k || nums[u-k] != v {
				nums[u] = v
				u++
			}
		}
		return u
	}
	return process(2)
}
