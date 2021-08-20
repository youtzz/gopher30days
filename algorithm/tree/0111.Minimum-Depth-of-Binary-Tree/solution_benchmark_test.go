package problem0111

import (
	"github.com/sevenger/gopher30days/algorithm"
	"testing"
)

/**
* goos: linux
* goarch: amd64
* pkg: github.com/sevenger/gopher30days/algorithm/tree/0111.Minimum-Depth-of-Binary-Tree
* cpu: Intel(R) Core(TM) i5-9400 CPU @ 2.90GHz
* Benchmark_minDepth-6       	20187448	        58.55 ns/op
* Benchmark_minDepth_DFS-6   	18622622	        63.70 ns/op
* PASS
批注：求最短路径BFS总是更快，因为BFS在找到target之后就停下了，而DFS需要递归整个树才能找到最短路径
*/

func Benchmark_minDepth(b *testing.B) {
	tree := algorithm.NewBinaryTreeByArgs(3, 9, 20, algorithm.NULL, algorithm.NULL, 15, 7)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		minDepth(tree)
	}
}

func Benchmark_minDepth_DFS(b *testing.B) {
	tree := algorithm.NewBinaryTreeByArgs(3, 9, 20, algorithm.NULL, algorithm.NULL, 15, 7)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		minDepth_DFS(tree)
	}
}
