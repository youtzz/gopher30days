package _54_Find_Minimum_in_Rotated_Sorted_Array_ii

func findMin(nums []int) int {
	n := len(nums)
	min := nums[0]
	left, right := 0, n-1
	for left <= right {
		mid := (right-left)>>1 + left
		if nums[left] == nums[mid] && nums[right] == nums[mid] {
			left++
			right--
		} else if nums[left] <= nums[mid] {
			if min > nums[left] {
				min = nums[left]
			}
			left = mid + 1
		} else {
			if min > nums[mid] {
				min = nums[mid]
			}
			right = mid - 1
		}
	}

	left--
	if left < n && nums[left] < min {
		min = nums[left]
	}
	return min
}
