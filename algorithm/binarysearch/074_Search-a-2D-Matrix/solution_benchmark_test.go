package searcha2dmatrix

import "testing"

func Benchmark_searchMatrix(b *testing.B) {
	args := tests[0].args
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		searchMatrix(args.matrix, args.target)
	}
}

func Benchmark_searchMatrix_best(b *testing.B) {
	args := tests[0].args
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		searchMatrix_best(args.matrix, args.target)
	}
}
