package binarysearch

// first solution
func search(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := (right-left)>>1 + left
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid
		} else {
			return mid
		}
	}
	return -1
}
