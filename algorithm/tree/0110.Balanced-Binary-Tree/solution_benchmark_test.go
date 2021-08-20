package problem0110

import (
	"github.com/sevenger/gopher30days/algorithm"
	"testing"
)

/**
 * goos: linux
 * goarch: amd64
 * pkg: github.com/sevenger/gopher30days/algorithm/tree/0110.Balanced-Binary-Tree
 * cpu: Intel(R) Core(TM) i5-9400 CPU @ 2.90GHz
 * Benchmark_isBalanced-6        	57098278	        20.89 ns/op
 * Benchmark_isBalanced_best-6   	60441040	        19.84 ns/op
 * PASS
 */

func Benchmark_isBalanced(b *testing.B) {
	tree := algorithm.NewBinaryTreeByArgs(3, 9, 20, algorithm.NULL, algorithm.NULL, 15, 7)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isBalanced(tree)
	}
}

func Benchmark_isBalanced_best(b *testing.B) {
	tree := algorithm.NewBinaryTreeByArgs(3, 9, 20, algorithm.NULL, algorithm.NULL, 15, 7)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isBalanced_best(tree)
	}
}
