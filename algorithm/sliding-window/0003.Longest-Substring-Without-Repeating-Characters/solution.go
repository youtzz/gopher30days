package problem0003

func lengthOfLongestSubstring(s string) int {
	window := make([]byte, 255)
	left, right := 0, 0
	max := 0
	for right < len(s) {
		in := s[right]
		right++
		window[in]++

		for window[in] > 1 {
			out := s[left]
			left++
			window[out]--
		}
		if right-left > max {
			max = right - left
		}
	}
	return max
}
