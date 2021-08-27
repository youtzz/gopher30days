package problem0309

func maxProfit(prices []int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}

	dp := make([][2]int, n)
	// 第一天买入
	dp[0][1] = -prices[0]
	// 第二天未持有：昨天未持有、昨天持有今天卖
	dp[1][0] = max(dp[0][0], dp[0][1]+prices[1])
	// 第二天持有：昨天持有、昨天未持有今天买
	dp[1][1] = max(dp[0][1], dp[0][0]-prices[1])

	for i := 2; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		// 今天持有：昨天就持有、前天未持有，今天买入
		dp[i][1] = max(dp[i-1][1], dp[i-2][0]-prices[i])
	}
	return dp[n-1][0]
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
