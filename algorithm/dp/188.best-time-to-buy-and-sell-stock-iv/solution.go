package problem0188

import "math"

const null = math.MinInt32

func maxProfit(K int, prices []int) int {
	n := len(prices)
	if n <= 1 || K <= 0 {
		return 0
	}

	// 天数、交易次数、持有状态
	dp := make([][][2]int, n)
	for i := 0; i < n; i++ {
		for k := 0; k <= K; k++ {
			dp[i] = append(dp[i], [2]int{})
		}
	}

	// base case
	// 第一天时，最多只有一次交易，其余都非法
	for i := 0; i <= K; i++ {
		dp[0][i][1] = null
	}
	dp[0][1][1] = -prices[0]

	// 记录所有交易时的最大利润
	var ans int
	for i := 1; i < len(prices); i++ {
		for k := 1; k <= K; k++ {
			dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
			dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
			ans = max(ans, dp[i][k][0])
		}
	}

	return ans
}

// 在我解的基础上进行优化
func maxProfit_best(K int, prices []int) int {
	n := len(prices)
	if n <= 1 || K <= 0 {
		return 0
	}
	// 优化点：一次交易包含买和卖，需要间隔两天，所以k应该小于n/2，超出相当于k=infinite，可以使用无k的动态规划写法
	if K > n/2 {
		return maxProfit_infinity(prices)
	}

	// 天数、交易次数、持有状态
	dp := make([][][2]int, n)
	for i := 0; i < n; i++ {
		for k := 0; k <= K; k++ {
			dp[i] = append(dp[i], [2]int{})
		}
	}

	// base case
	for i := 0; i <= K; i++ {
		dp[0][i][1] = null
	}
	dp[0][1][1] = -prices[0]

	for i := 1; i < len(prices); i++ {
		for k := 1; k <= K; k++ {
			dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
			dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
		}
	}

	ans := 0
	for k := 1; k <= K; k++ {
		ans = max(dp[len(prices)-1][k][0], ans)
	}
	return ans
}

func maxProfit_infinity(price []int) int {
	dp_i_0, dp_i_1 := 0, -price[0]

	for i := 1; i < len(price); i++ {
		dp_i_0, dp_i_1 = max(dp_i_0, dp_i_1+price[i]), max(dp_i_1, dp_i_0-price[i])
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
