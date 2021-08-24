package problem0438

func findAnagrams(s string, p string) []int {
	need, window := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(p); i++ {
		need[p[i]]++
	}

	var res []int
	left, right := 0, 0
	valid := 0
	for right < len(s) {
		in := s[right]
		right++

		if _, has := need[in]; has {
			window[in]++
			if window[in] == need[in] {
				valid++
			}
		}

		for right-left >= len(p) {
			if valid == len(need) {
				res = append(res, left)
			}
			out := s[left]
			left++
			if _, has := need[out]; has {
				if window[out] == need[out] {
					valid--
				}
				window[out]--
			}
		}
	}

	return res
}
