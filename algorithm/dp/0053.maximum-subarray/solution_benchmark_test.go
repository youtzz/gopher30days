package problem0053

import "testing"

/**
* goos: linux
* goarch: amd64
* pkg: github.com/sevenger/gopher30days/algorithm/dp/0053.maximum-subarray
* cpu: Intel(R) Core(TM) i5-9400 CPU @ 2.90GHz
* Benchmark_maxSubArray-6      	38268223	        31.43 ns/op
* Benchmark_maxSubArray_dp-6   	32472376	        43.10 ns/op
* PASS
 */

func Benchmark_maxSubArray(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	}
}

func Benchmark_maxSubArray_dp(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maxSubArray_dp([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	}
}
