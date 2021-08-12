package searchRange

// first solution
func searchRange(nums []int, target int) []int {
	res := make([]int, 2)
	res[0], res[1] = -1, -1
	if len(nums) == 0 {
		return res
	}

	// 第一次二分搜索，查找第一次出现target的位置
	left, right := 0, len(nums)-1
	for left+1 < right {
		mid := (right-left)>>1 + left
		if nums[mid] < target {
			left = mid
		} else if nums[mid] > target {
			right = mid
		} else {
			right = mid
		}
	}
	if nums[left] == target {
		res[0] = left
	} else if nums[right] == target {
		res[0] = right
	} else {
		return res
	}

	// 第二次二分搜索，查找第二次出现target的位置
	left, right = 0, len(nums)-1
	for left+1 < right {
		mid := (right-left)>>1 + left
		if nums[mid] < target {
			left = mid
		} else if nums[mid] > target {
			right = mid
		} else {
			left = mid
		}
	}
	if nums[right] == target {
		res[1] = right
	} else if nums[left] == target {
		res[1] = left
	}

	return res
}
