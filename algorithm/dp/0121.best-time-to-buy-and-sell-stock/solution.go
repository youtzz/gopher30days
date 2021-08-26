package problem0121

func maxProfit(prices []int) int {
	// 天数、持有状态
	dp := make([][2]int, len(prices))
	// base case
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	for i := 1; i < len(dp); i++ {
		// 今天没有持有股票，是因为：1、昨天就没有，今天继续没有 2、昨天有，今天卖了
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		// 今天持有股票，是因为：1、昨天就有，今天继续有 2、昨天没有的，今天买了
		dp[i][1] = max(dp[i-1][1], -prices[i])
	}

	return dp[len(prices)-1][0]
}

// 因为这题的状态只跟相邻状态有关，所以可以不用dp数组，直接变量保存
func maxProfit_best(prices []int) int {
	// 天数、持有状态
	var dp_i_0 int // 第i天未持有的收益
	var dp_i_1 int // 第i持有的收益

	// base case
	dp_i_0 = 0
	dp_i_1 = -prices[0]
	for i := 1; i < len(prices); i++ {
		// 今天没有持有股票，是因为：1、昨天就没有，今天继续没有 2、昨天有，今天卖了
		dp_i_0 = max(dp_i_0, dp_i_1+prices[i])
		// 今天持有股票，是因为：1、昨天就有，今天继续有 2、昨天没有的，今天买了
		dp_i_1 = max(dp_i_1, -prices[i])
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
