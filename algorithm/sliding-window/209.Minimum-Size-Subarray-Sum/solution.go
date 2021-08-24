package problem0209

import "math"

func minSubArrayLen(target int, nums []int) int {
	left, right := 0, 0
	sum := 0
	ans := math.MaxInt32
	for right < len(nums) {
		// 窗口不断右移
		in := nums[right]
		right++

		sum += in
		// 当sum>=target时窗口左移
		for sum >= target {
			if right-left < ans {
				ans = right - left
			}

			out := nums[left]
			left++
			sum -= out
		}
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}
