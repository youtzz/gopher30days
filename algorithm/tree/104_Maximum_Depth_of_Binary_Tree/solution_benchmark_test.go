package Maximum_Depth_of_Binary_Tree

import (
	. "github.com/sevenger/gopher30days/algorithm"
	"testing"
)

func BenchmarkMaxDepth_bfs(b *testing.B) {
	tree := NewBinaryTreeByArgs("3", "9", "20", "nil", "nil", "15", "7")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maxDepth_bfs(tree)
	}
}

func BenchmarkMaxDepth(b *testing.B) {
	tree := NewBinaryTreeByArgs("3", "9", "20", "nil", "nil", "15", "7")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maxDepth(tree)
	}
}
