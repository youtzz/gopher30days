package problem0283

func moveZeroes(nums []int) {
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[index], nums[i] = nums[i], nums[index]
			index++
		}
	}
}
