package problem94

import "testing"

var arg = tests[0].args

func Benchmark_inorderTraversal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		inorderTraversal(arg.root)
	}
}

func Benchmark_inorderTraversal_interactive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		inorderTraversal_interactive(arg.root)
	}
}
