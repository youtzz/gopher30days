package problem0026

func removeDuplicates(nums []int) int {
	index := 0
	for i := 0; i < len(nums); i++ {
		// 找到重复的num的最后一个值
		num := nums[i]
		for i < len(nums) && nums[i] == num {
			i++
		}
		i--

		nums[index] = nums[i]
		index++
	}
	return index
}

func removeDuplicates_best(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slow := 1
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
