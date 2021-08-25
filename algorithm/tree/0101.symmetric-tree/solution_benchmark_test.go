package problem0101

import (
	"github.com/sevenger/gopher30days/algorithm"
	"testing"
)

/**
 * goos: linux
 * goarch: amd64
 * pkg: github.com/sevenger/gopher30days/algorithm/tree/0101.symmetric-tree
 * cpu: Intel(R) Core(TM) i5-9400 CPU @ 2.90GHz
 * Benchmark_isSymmetric
 * Benchmark_isSymmetric-6              	 2063042	       574.0 ns/op
 * Benchmark_isSymmetric_best
 * Benchmark_isSymmetric_best-6         	 3023816	       453.6 ns/op
 * Benchmark_isSymmetric_best_recur
 * Benchmark_isSymmetric_best_recur-6   	35288827	        30.37 ns/op
 * PASS
 */

func Benchmark_isSymmetric(b *testing.B) {
	tree := algorithm.NewBinaryTreeByArgs(1, 2, 2, 3, 4, 4, 3)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isSymmetric(tree)
	}
}

func Benchmark_isSymmetric_best(b *testing.B) {
	tree := algorithm.NewBinaryTreeByArgs(1, 2, 2, 3, 4, 4, 3)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isSymmetric_best(tree)
	}
}

func Benchmark_isSymmetric_best_recur(b *testing.B) {
	tree := algorithm.NewBinaryTreeByArgs(1, 2, 2, 3, 4, 4, 3)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isSymmetric_best_recur(tree)
	}
}
