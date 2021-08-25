package problem0070

import "testing"

/*
 * goos: linux
 * goarch: amd64
 * pkg: github.com/sevenger/gopher30days/algorithm/dp/70.climbing-stairs
 * cpu: Intel(R) Core(TM) i5-9400 CPU @ 2.90GHz
 * Benchmark_climbStairs
 * Benchmark_climbStairs-6        	 4158352	       267.2 ns/op
 * Benchmark_climbStairs_best
 * Benchmark_climbStairs_best-6   	31199920	        37.08 ns/op
 * PASS
 */

// dp因为有dp[]，有更多的空间操作，速度更慢
func Benchmark_climbStairs(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		climbStairs(100)
	}
}

func Benchmark_climbStairs_best(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		climbStairs_best(100)
	}
}
