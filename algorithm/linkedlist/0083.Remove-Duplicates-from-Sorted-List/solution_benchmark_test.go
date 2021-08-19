package problem0083

import (
	"github.com/sevenger/gopher30days/algorithm"
	"testing"
)

func Benchmark_deleteDuplicates(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		deleteDuplicates(algorithm.NewLinkedListByArgs(1, 1, 2, 2, 2, 2, 2, 2, 2))
	}
}

func Benchmark_deleteDuplicates2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		deleteDuplicates2(algorithm.NewLinkedListByArgs(1, 1, 2, 2, 2, 2, 2, 2, 2))
	}
}

func Benchmark_deleteDuplicates_recur(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		deleteDuplicates_recur(algorithm.NewLinkedListByArgs(1, 1, 2, 2, 2, 2, 2, 2, 2))
	}
}

func Benchmark_deleteDuplicates_best(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		deleteDuplicates_best(algorithm.NewLinkedListByArgs(1, 1, 2, 2, 2, 2, 2, 2, 2))
	}
}
