package problem0104

import (
	"github.com/sevenger/gopher30days/algorithm"
	"testing"
)

func BenchmarkMaxDepth_bfs(b *testing.B) {
	tree := algorithm.NewBinaryTreeByArgs("3", "9", "20", "nil", "nil", "15", "7")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maxDepth_bfs(tree)
	}
}

func BenchmarkMaxDepth(b *testing.B) {
	tree := algorithm.NewBinaryTreeByArgs("3", "9", "20", "nil", "nil", "15", "7")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maxDepth(tree)
	}
}
