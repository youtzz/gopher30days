package problem0188

import "testing"

/**
 * goos: linux
 * goarch: amd64
 * pkg: github.com/sevenger/gopher30days/algorithm/dp/188.best-time-to-buy-and-sell-stock-iv
 * cpu: Intel(R) Core(TM) i5-9400 CPU @ 2.90GHz
 * Benchmark_maxProfit
 * Benchmark_maxProfit-6        	 1180131	      1169 ns/op
 * Benchmark_maxProfit_best
 * Benchmark_maxProfit_best-6   	186173425	         6.432 ns/op
 * PASS
批注：在 k>len(prices)/2 时 best 方法速度激增
*/

func Benchmark_maxProfit(b *testing.B) {
	k := 4
	prices := []int{3, 2, 6, 5, 0, 3}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maxProfit(k, prices)
	}
}

func Benchmark_maxProfit_best(b *testing.B) {
	k := 4
	prices := []int{3, 2, 6, 5, 0, 3}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maxProfit_best(k, prices)
	}
}
