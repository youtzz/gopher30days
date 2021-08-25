package problem0169

func majorityElement(nums []int) int {
	set := make(map[int]int, 9)
	for i := 0; i < len(nums); i++ {
		set[nums[i]]++
	}
	ans, max := 0, 0
	for k, v := range set {
		if v > max {
			ans, max = k, v
		}
	}
	return ans
}
