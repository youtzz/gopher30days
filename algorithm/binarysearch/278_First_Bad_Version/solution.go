package firstbadversion

import "sort"

// first solution
func firstBadVersion(n int) int {
	left, right := 1, n
	for left < right {
		mid := (right + left) / 2
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

// best solution
func firstBadVersion_best(n int) int {
	return sort.Search(n, func(version int) bool {
		return isBadVersion(version)
	})
}

func isBadVersion(n int) (res bool) {
	if n >= 5 {
		res = true
	}
	return
}
