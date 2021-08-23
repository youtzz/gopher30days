package problem0509

import "testing"

/**
 * goos: linux
 * goarch: amd64
 * pkg: github.com/sevenger/gopher30days/algorithm/dp/0509.Fibonacci-Number
 * cpu: Intel(R) Core(TM) i5-9400 CPU @ 2.90GHz
 * BenchmarkFib_Recur-6   	 4788055	       247.6 ns/op
 * BenchmarkFib_mem-6     	17869836	        96.72 ns/op
 * BenchmarkFib_dp-6      	20777540	        59.73 ns/op
 * PASS
 */

// 传统递归法，最慢
func BenchmarkFib_Recur(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fib_recur(10)
	}
}

// 在传统递归法上使用了备忘录，对递归树进行了剪枝
func BenchmarkFib_mem(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fib_mem(10)
	}
}

// dp法，最快
func BenchmarkFib_dp(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fib_dp(10)
	}
}
