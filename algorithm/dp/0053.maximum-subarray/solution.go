package problem0053

import "math"

// 暴力解
func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	max := math.MinInt32
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum > max {
				max = sum
			}
		}
	}
	return max
}

// 动态规划
func maxSubArray_dp(nums []int) int {
	var max = func(x, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}

	// dp[i]代表以nums[i]结尾的最大子序和
	dp := make([]int, len(nums))

	// base case
	dp[0] = nums[0]

	ans := dp[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		ans = max(ans, dp[i])
	}
	return ans
}
