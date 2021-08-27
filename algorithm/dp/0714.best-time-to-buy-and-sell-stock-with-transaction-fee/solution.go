package problem0714

func maxProfit(prices []int, fee int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}

	dp := make([][2]int, n)

	// 买股票时收手续费
	dp[0][1] = -prices[0] - fee

	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i]-fee)
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
