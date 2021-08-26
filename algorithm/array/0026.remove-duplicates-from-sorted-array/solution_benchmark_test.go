package problem0026

import "testing"

/**
 * goos: linux
 * goarch: amd64
 * pkg: github.com/sevenger/gopher30days/algorithm/array/0026.remove-duplicates-from-sorted-array
 * cpu: Intel(R) Core(TM) i5-9400 CPU @ 2.90GHz
 * Benchmark_removeDuplicates
 * Benchmark_removeDuplicates-6        	39449762	        30.46 ns/op
 * Benchmark_removeDuplicates_best
 * Benchmark_removeDuplicates_best-6   	196164670	         6.192 ns/op
 * PASS
 */

func Benchmark_removeDuplicates(b *testing.B) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		removeDuplicates(nums)
	}
}

func Benchmark_removeDuplicates_best(b *testing.B) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		removeDuplicates_best(nums)
	}
}
