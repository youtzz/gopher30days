package algorithm

import "testing"

func BenchmarkMergeSort(b *testing.B) {
	arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MergeSort(arr)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		QuickSort(arr)
	}
}

func BenchmarkDfs(b *testing.B) {
	tree := NewBinaryTree([]string{"1", "2", "3", "4", "5", "6", "7"})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Dfs(tree)
	}
}

func BenchmarkDfs_divideAndConquer(b *testing.B) {
	tree := NewBinaryTree([]string{"1", "2", "3", "4", "5", "6", "7"})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Dfs_divideAndConquer(tree)
	}
}
