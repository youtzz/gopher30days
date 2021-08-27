package problem0122

func maxProfit(prices []int) int {
	n := len(prices)
	dp := make([][2]int, n)

	// base case
	dp[0][1] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}

	return dp[n-1][0]
}

// 最优解，简化dp数组
func maxProfit_best(prices []int) int {
	var dp_i_0, dp_i_1 int

	// base case
	dp_i_1 = -prices[0]

	for i := 1; i < len(prices); i++ {
		dp_i_0, dp_i_1 = max(dp_i_0, dp_i_1+prices[i]), max(dp_i_1, dp_i_0-prices[i])
	}

	return dp_i_0
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
