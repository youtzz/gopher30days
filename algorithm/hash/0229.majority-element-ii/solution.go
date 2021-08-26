package problem0229

func majorityElement(nums []int) []int {
	set := make(map[int]int)
	res := make([]int, 0)
	l := len(nums) / 3
	for i := 0; i < len(nums); i++ {
		set[nums[i]]++
		if set[nums[i]]-1 == l {
			res = append(res, nums[i])
		}
	}
	return res
}
