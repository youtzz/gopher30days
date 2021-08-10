package _81_Search_in_Rotated_Sorted_Array_ii

func search(nums []int, target int) bool {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right-left)>>1 + left
		if nums[mid] == target {
			return true
		}
		// left和right相等时分别右移和左移
		if nums[left] == nums[mid] && nums[right] == nums[mid] {
			left++
			right--
		} else if nums[left] <= nums[mid] {
			if nums[left] <= target && nums[mid] > target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[right] >= target && nums[mid] < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return false
}
