package problem0209

import "math"

// 第一次提交的代码，可以算出答案，但是超时了
func maxSlidingWindow(nums []int, k int) []int {
	ans := make([]int, 0, len(nums)/k)
	num := math.MinInt32
	left, right := 0, 0
	for right < len(nums) {
		num = max(num, nums[right])
		right++

		if right-left == k {
			ans = append(ans, num)
			num = math.MinInt32
			left++
			right = left
		}
	}
	return ans
}

// 最优解
func maxSlidingWindow_best(nums []int, k int) []int {
	q := []int{}
	push := func(i int) {
		for len(q) > 0 && nums[i] >= nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}

	for i := 0; i < k; i++ {
		push(i)
	}

	n := len(nums)
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[q[0]]
	for i := k; i < n; i++ {
		push(i)
		for q[0] <= i-k {
			q = q[1:]
		}
		ans = append(ans, nums[q[0]])
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
