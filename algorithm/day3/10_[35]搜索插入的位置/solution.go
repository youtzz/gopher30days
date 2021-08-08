package main

// first solution
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := (right + left) / 2
		n := nums[mid]
		if n < target {
			left = mid + 1
		} else if n > target {
			right = mid
		} else {
			return mid
		}
	}
	return left
}
