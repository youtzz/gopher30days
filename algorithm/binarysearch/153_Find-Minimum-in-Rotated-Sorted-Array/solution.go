package _53_Find_Minimum_in_Rotated_Sorted_Array

func findMin(nums []int) int {
	min := nums[0]
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right-left)>>1 + left
		if nums[left] <= nums[mid] {
			if nums[left] < min {
				min = nums[left]
			}
			left = mid + 1
		} else {
			if nums[mid] < min {
				min = nums[mid]
			}
			right = mid - 1
		}
	}

	return min
}

func findMin_best(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		pivot := low + (high-low)/2
		if nums[pivot] < nums[high] {
			high = pivot
		} else {
			low = pivot + 1
		}
	}
	return nums[low]
}
