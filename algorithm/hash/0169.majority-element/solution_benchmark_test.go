package problem0169

import "testing"

/**
 * goos: linux
 * goarch: amd64
 * pkg: github.com/sevenger/gopher30days/algorithm/hash/0169.majority-element
 * cpu: Intel(R) Core(TM) i5-9400 CPU @ 2.90GHz
 * Benchmark_majorityElement
 * Benchmark_majorityElement-6    	 3717650	       356.0 ns/op
 * Benchmark_majorityElement2
 * Benchmark_majorityElement2-6   	186010255	         6.431 ns/op
 * PASS
 */

func Benchmark_majorityElement(b *testing.B) {
	arr := []int{-1, 100, 2, 100, 100, 4, 100}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		majorityElement(arr)
	}
}

func Benchmark_majorityElement2(b *testing.B) {
	arr := []int{-1, 100, 2, 100, 100, 4, 100}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		majorityElement2(arr)
	}
}
