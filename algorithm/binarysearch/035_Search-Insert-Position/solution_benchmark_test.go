package searchInsertPosition

import "testing"

func Benchmark_searchInsert(b *testing.B) {
	args := tests[0].args
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		searchInsert(args.nums, args.target)
	}
}

func Benchmark_searchInsert2(b *testing.B) {
	args := tests[0].args
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		searchInsert2(args.nums, args.target)
	}
}
