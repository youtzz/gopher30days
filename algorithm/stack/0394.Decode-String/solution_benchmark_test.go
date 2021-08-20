package problem0394

import "testing"

const str = "3[a2[c]]"

func Benchmark_decodeString(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		decodeString(str)
	}
}

func Benchmark_decodeString_best(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		decodeString(str)
	}
}
