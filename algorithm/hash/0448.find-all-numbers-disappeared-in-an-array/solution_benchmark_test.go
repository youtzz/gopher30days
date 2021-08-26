package problem0448

import "testing"

/**
 * goos: linux
 * goarch: amd64
 * pkg: github.com/sevenger/gopher30days/algorithm/array/0448.find-all-numbers-disappeared-in-an-array
 * cpu: Intel(R) Core(TM) i5-9400 CPU @ 2.90GHz
 * Benchmark_findDisappearedNumbers
 * Benchmark_findDisappearedNumbers-6        	13357116	        90.08 ns/op
 * Benchmark_findDisappearedNumbers2
 * Benchmark_findDisappearedNumbers2-6       	24375800	        50.90 ns/op
 * Benchmark_findDisappearedNumbers_best
 * Benchmark_findDisappearedNumbers_best-6   	 8321118	       142.5 ns/op
 * PASS
 */

func Benchmark_findDisappearedNumbers(b *testing.B) {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	b.ResetTimer()
	for i := 1; i < b.N; i++ {
		findDisappearedNumbers(nums)
	}
}

func Benchmark_findDisappearedNumbers2(b *testing.B) {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findDisappearedNumbers2(nums)
	}
}

func Benchmark_findDisappearedNumbers_best(b *testing.B) {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findDisappearedNumbers_best(nums)
	}
}
