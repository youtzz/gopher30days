package problem0300

func lengthOfLIS(nums []int) int {
	// dp 含义: dp[i]表示nums[i]这个数为结尾的最长递增子序列长度
	dp := make([]int, len(nums))
	// base case: 所有值都是1
	for i, _ := range dp {
		dp[i] = 1
	}

	for i, _ := range dp {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	res := 0
	for _, v := range dp {
		if v > res {
			res = v
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
