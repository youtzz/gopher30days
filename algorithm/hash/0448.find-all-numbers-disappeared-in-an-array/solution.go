package problem0448

// hash法
func findDisappearedNumbers(nums []int) []int {
	var res []int

	// 存储 [1, len(nums)]出现的次数
	need := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if num >= 1 && num <= len(nums) {
			need[num]++
		}
	}

	// 所有出现0次的数就是消失的数字
	for i := 1; i < len(need); i++ {
		if need[i] == 0 {
			res = append(res, i)
		}
	}

	return res
}

// 第二次解，空间复杂度为 O(1)（返回的切片不算）
func findDisappearedNumbers2(nums []int) []int {
	// 将[1,n]加入res
	res := make([]int, len(nums)+1)
	for i := 0; i <= len(nums); i++ {
		res[i] = i
	}

	// nums中出现的数字在res中标记为0
	for _, v := range nums {
		if 1 <= v && v <= len(nums) {
			res[v] = 0
		}
	}

	// 去除res中标记为0的值
	index := 0
	for i := 0; i < len(res); i++ {
		if res[i] != 0 {
			res[index], res[i] = res[i], res[index]
			index++
		}
	}
	return res[:index]
}

// best解
func findDisappearedNumbers_best(nums []int) (ans []int) {
	n := len(nums)
	for _, v := range nums {
		v = (v - 1) % n
		nums[v] += n
	}
	for i, v := range nums {
		if v <= n {
			ans = append(ans, i+1)
		}
	}
	return
}
