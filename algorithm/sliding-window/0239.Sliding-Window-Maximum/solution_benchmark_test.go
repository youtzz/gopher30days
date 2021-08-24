package problem0209

import "testing"

/**
 * goos: linux
 * goarch: amd64
 * pkg: github.com/sevenger/gopher30days/algorithm/sliding-window/0239.Sliding-Window-Maximum
 * cpu: Intel(R) Core(TM) i5-9400 CPU @ 2.90GHz
 * BenchmarkMaxSlidingWindow-6       	   17220	     87931 ns/op
 * BenchmarkMaxSlidingWindowBest-6   	   28782	     45709 ns/op
 * PASS
批注：在数据量很大时，best写法更快，然而在小数据量时，我的写法更快。
*/

func BenchmarkMaxSlidingWindow(b *testing.B) {
	var arr []int
	for i := 0; i < 10000; i++ {
		arr = append(arr, i)
	}
	//arr = []int{1, 3, -1, -3, 5, 3, 6, 7}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maxSlidingWindow(arr, 3)
	}
}

func BenchmarkMaxSlidingWindowBest(b *testing.B) {
	var arr []int
	for i := 0; i < 10000; i++ {
		arr = append(arr, i)
	}
	//arr = []int{1, 3, -1, -3, 5, 3, 6, 7}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maxSlidingWindow_best(arr, 3)
	}
}
