package main

import "sort"

// first solution
func searchMatrix(matrix [][]int, target int) bool {
	row := sort.Search(len(matrix), func(i int) bool { return matrix[i][0] > target }) - 1
	col := sort.SearchInts(matrix[row], target)
	return col < len(matrix[row]) && matrix[row][col] == target
}

// best solution
func searchMatrix_best(matrix [][]int, target int) bool {
	m, n := len(matrix[0]), len(matrix)
	left, right := 0, m*n-1
	for left <= right {
		mid := (right-left)>>1 + left
		num := matrix[mid/m][mid%m]
		if num == target {
			return true
		} else if num > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return false
}
