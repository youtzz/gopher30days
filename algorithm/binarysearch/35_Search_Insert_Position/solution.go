package searchInsertPosition

import "sort"

// first solution
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := (right-left)>>1 + left
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func searchInsert2(nums []int, target int) int {
	return sort.Search(len(nums), func(i int) bool {
		return nums[i] >= target
	})
}
