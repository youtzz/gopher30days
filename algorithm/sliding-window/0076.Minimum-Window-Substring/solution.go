package problem0076

import "math"

func minWindow(s string, t string) string {
	need, window := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	left, right := 0, 0
	valid := 0

	start, length := 0, math.MaxInt32
	for right < len(s) {
		c := s[right]
		right++

		if _, has := need[c]; has {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		for valid == len(need) {
			if right-left < length {
				start = left
				length = right - left
			}

			d := s[left]
			left++
			if _, has := need[d]; has {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}

	if length == math.MaxInt32 {
		return ""
	} else {
		return s[start : start+length]
	}
}
