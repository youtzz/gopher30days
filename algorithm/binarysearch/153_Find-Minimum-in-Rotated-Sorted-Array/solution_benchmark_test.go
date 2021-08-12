package _53_Find_Minimum_in_Rotated_Sorted_Array

import "testing"

func Benchmark_findMin(b *testing.B) {
	nums := tests[0].args.nums
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findMin(nums)
	}
}

func Benchmark_findMin_best(b *testing.B) {
	nums := tests[0].args.nums
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findMin_best(nums)
	}
}
