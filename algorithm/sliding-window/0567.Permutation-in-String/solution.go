package problem0567

func checkInclusion(s1 string, s2 string) bool {
	need, window := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		need[s1[i]]++
	}

	left, right := 0, 0
	valid := 0
	for right < len(s2) {
		in := s2[right]
		right++

		if _, has := need[in]; has {
			window[in]++
			if window[in] == need[in] {
				valid++
			}
		}
		for right-left >= len(s1) {
			if valid == len(need) {
				return true
			}
			out := s2[left]
			left++
			if _, has := need[out]; has {
				if window[out] == need[out] {
					valid--
				}
				window[out]--
			}
		}
	}
	return false
}
