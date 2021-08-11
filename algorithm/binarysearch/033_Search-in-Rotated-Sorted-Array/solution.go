package _33_Search_in_Rotated_Sorted_Array

// first solution
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right-left)>>1 + left
		if nums[mid] == target {
			return mid
		}
		if nums[0] <= nums[mid] {
			// 左半边是有序的
			if nums[left] <= target && nums[mid] >= target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// 右半边是有序的
			if nums[right] >= target && nums[mid] <= target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
