package searcha2dmatrix

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
	index := sort.Search(m*n, func(i int) bool { return matrix[i/m][i%m] >= target })
	return index < m*n && matrix[index/m][index%m] == target
}
