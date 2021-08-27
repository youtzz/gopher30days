package problem0123

import "math"

func maxProfit(prices []int) int {
	n := len(prices)
	// 天数、卖过几次（交易次数）、持有状态
	dp := make([][3][2]int, n)

	const null = math.MinInt32

	// base case
	// 第一天不买和买
	dp[0][0][0], dp[0][0][1] = 0, -prices[0]
	// 第一天无法交易
	dp[0][1][0], dp[0][1][1] = null, null
	dp[0][2][0], dp[0][2][1] = null, null

	for i := 1; i < n; i++ {
		// 卖过0次，未持有
		dp[i][0][0] = 0
		// 卖过0次，持有:昨天就持有、昨天未持有，今天买
		dp[i][0][1] = max(dp[i-1][0][1], dp[i-1][0][0]-prices[i])
		// 卖过1次，未持有：昨天就未持有、昨天持有，今天卖
		dp[i][1][0] = max(dp[i-1][1][0], dp[i-1][0][1]+prices[i])
		// 卖过1次，持有：昨天就持有、昨天未持有，今天买
		dp[i][1][1] = max(dp[i-1][1][1], dp[i-1][1][0]-prices[i])
		// 卖过两次，未持有：昨天就未持有、昨天持有，今天卖
		dp[i][2][0] = max(dp[i-1][2][0], dp[i-1][1][1]+prices[i])
		// 卖过两次，持有：不可能
		dp[i][2][1] = null
	}

	// 比较最后一天，交易两次、交易一次、不交易的利润
	return max(max(dp[n-1][2][0], dp[n-1][1][0]), dp[n-1][0][0])
}

// 买入记为交易一次
func maxProfit2(prices []int) int {
	const null = math.MinInt32
	const MaxK = 2
	n := len(prices)
	// 天数、卖过几次（交易次数）、持有状态
	dp := make([][3][2]int, n)
	// base case
	// 第一天未买入却持有：不可能
	dp[0][0][1] = null
	// 第一天买入却未持有：不可能
	dp[0][1][0] = null
	// 第一天买入且持有
	dp[0][1][1] = -prices[0]
	// 第一天买入两次：不可能
	dp[0][2][0], dp[0][2][1] = null, null

	for i := 1; i < n; i++ {
		for k := 1; k <= MaxK; k++ {
			dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
			dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
		}
	}

	// 比较最后一天，交易两次、交易一次、不交易的利润
	return max(max(dp[n-1][2][0], dp[n-1][1][0]), 0)
}

// 最优解：简化dp数组
func maxProfit_best(prices []int) int {
	const null = math.MinInt32
	n := len(prices)
	// 天数、卖过几次（交易次数）、持有状态
	// base case
	dp_i_0_1 := -prices[0]
	dp_i_1_0, dp_i_1_1 := null, null
	dp_i_2_0 := null

	for i := 1; i < n; i++ {
		// 卖过0次，持有:昨天就持有、昨天未持有，今天买
		dp_i_0_1 = max(dp_i_0_1, -prices[i])
		// 卖过1次，未持有：昨天就未持有、昨天持有，今天卖
		dp_i_1_0 = max(dp_i_1_0, dp_i_0_1+prices[i])
		// 卖过1次，持有：昨天就持有、昨天未持有，今天买
		dp_i_1_1 = max(dp_i_1_1, dp_i_1_0-prices[i])
		// 卖过两次，未持有：昨天就未持有、昨天持有，今天卖
		dp_i_2_0 = max(dp_i_2_0, dp_i_1_1+prices[i])
	}

	// 比较最后一天，交易两次、交易一次、不交易的利润
	return max(max(dp_i_2_0, dp_i_1_0), 0)
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
