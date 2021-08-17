package Open_the_Lock

import "testing"

/**
* goos: linux
* goarch: amd64
* pkg: github.com/sevenger/gopher30days/algorithm/tree/111_Minimum_Depth_of_Binary_Tree
* cpu: Intel(R) Core(TM) i5-9400 CPU @ 2.90GHz
* Benchmark_openLock-6                        2191            546059 ns/op          285418 B/op       6616 allocs/op
* Benchmark_openLock_double_bfs-6            15544             75739 ns/op           26755 B/op        687 allocs/op
* PASS
批注：双向BFS是单向BFS的10倍
*/

func Benchmark_openLock(b *testing.B) {
	deadends := []string{"0201", "0101", "0102", "1212", "2002"}
	target := "0202"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		openLock(deadends, target)
	}
}

func Benchmark_openLock_double_bfs(b *testing.B) {
	deadends := []string{"0201", "0101", "0102", "1212", "2002"}
	target := "0202"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		openLock_double_bfs(deadends, target)
	}
}
