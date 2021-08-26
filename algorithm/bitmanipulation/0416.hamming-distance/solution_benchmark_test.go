package problem0416

import "testing"

/**
 * goos: linux
 * goarch: amd64
 * pkg: github.com/sevenger/gopher30days/algorithm/bitmanipulation/0416.hamming-distance
 * cpu: Intel(R) Core(TM) i5-9400 CPU @ 2.90GHz
 * Benchmark_hammingDistance
 * Benchmark_hammingDistance-6        	125731779	         9.560 ns/op
 * Benchmark_hammingDistance2
 * Benchmark_hammingDistance2-6       	1000000000	         0.5181 ns/op
 * Benchmark_hammingDistance_best
 * Benchmark_hammingDistance_best-6   	661219412	         1.806 ns/op
 * PASS
 */

func Benchmark_hammingDistance(b *testing.B) {
	x, y := 1, 4
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hammingDistance(x, y)
	}
}

func Benchmark_hammingDistance2(b *testing.B) {
	x, y := 1, 4
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hammingDistance2(x, y)
	}
}

func Benchmark_hammingDistance_best(b *testing.B) {
	x, y := 1, 4
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hammingDistance_best(x, y)
	}
}
