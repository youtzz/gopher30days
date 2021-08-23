package problem0034

func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	if len(nums) == 0 {
		return res
	}

	// 寻找第一个出现的target的位置
	low, high := 0, len(nums)-1
	for low <= high {
		mid := (high-low)>>1 + low
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			// 相等时，如果mid是第一个||mid前面的数不等于target
			if mid == 0 || nums[mid-1] != target {
				res[0] = mid
				break
			} else {
				high = mid - 1
			}
		}
	}

	// 寻找最后一个出现的target的位置
	low, high = 0, len(nums)-1
	for low <= high {
		mid := (high-low)>>1 + low
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			// 相等时，如果mid是最后一个||mid后面的数不等于target
			if mid+1 == len(nums) || nums[mid+1] != target {
				res[1] = mid
				break
			} else {
				low = mid + 1
			}
		}
	}
	return res
}
