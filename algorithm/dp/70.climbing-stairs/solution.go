package problem0070

// 第一次解(终于能写出来动态规划了)
func climbStairs(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[1], dp[2] = 1, 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// 使用了滑动数组
func climbStairs_best(n int) int {
	if n == 1 || n == 2 {
		return n
	}

	p, q, r := 0, 0, 1
	for i := 0; i < n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}
