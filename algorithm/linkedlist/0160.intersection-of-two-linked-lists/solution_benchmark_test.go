package problem0160

import (
	"github.com/sevenger/gopher30days/algorithm"
	"testing"
)

/**
 * goos: linux
 * goarch: amd64
 * pkg: github.com/sevenger/gopher30days/algorithm/linkedlist/0160.intersection-of-two-linked-lists
 * cpu: Intel(R) Core(TM) i5-9400 CPU @ 2.90GHz
 * Benchmark_getIntersectionNode
 * Benchmark_getIntersectionNode-6        	13433571	        88.71 ns/op
 * Benchmark_getIntersectionNode2
 * Benchmark_getIntersectionNode2-6       	100000000	        11.30 ns/op
 * Benchmark_getIntersectionNode_best
 * Benchmark_getIntersectionNode_best-6   	100000000	        11.57 ns/op
 * PASS
 */

// 用了hash，所以很慢
func Benchmark_getIntersectionNode(b *testing.B) {
	headA, headB, _ := getUseCase(1)
	headA = algorithm.NewLinkedList(func() []int { return make([]int, 114514) }())
	headB = algorithm.NewLinkedList(func() []int { return make([]int, 10000) }())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		getIntersectionNode(headA, headB)
	}
}

func Benchmark_getIntersectionNode2(b *testing.B) {
	headA, headB, _ := getUseCase(1)
	headA = algorithm.NewLinkedList(func() []int { return make([]int, 114514) }())
	headB = algorithm.NewLinkedList(func() []int { return make([]int, 10000) }())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		getIntersectionNode2(headA, headB)
	}
}

func Benchmark_getIntersectionNode_best(b *testing.B) {
	headA, headB, _ := getUseCase(1)
	headA = algorithm.NewLinkedList(func() []int { return make([]int, 114514) }())
	headB = algorithm.NewLinkedList(func() []int { return make([]int, 10000) }())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		getIntersectionNode_best(headA, headB)
	}
}
