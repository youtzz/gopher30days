package problem0509

// 最简单暴力的解法，递归
func fib_recur(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fib_recur(n-1) + fib_recur(n-2)
}

// 使用了备忘录的改良法
var mem []int

func fib_mem(n int) int {
	mem = make([]int, n+1)
	return fibHelper(n)
}

func fibHelper(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	if mem[n] != 0 {
		return mem[n]
	}
	mem[n] = fibHelper(n-1) + fibHelper(n-2)
	return mem[n]
}

// dp法
func fib_dp(n int) int {
	if n == 0 {
		return 0
	}
	dp := make([]int, n+1)
	dp[0], dp[1] = 0, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
