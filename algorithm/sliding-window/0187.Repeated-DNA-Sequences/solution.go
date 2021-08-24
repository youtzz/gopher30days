package problem0187

func findRepeatedDnaSequences(s string) []string {
	window := make([]byte, 0)
	res := make(map[string]int)
	ans := make([]string, 0)

	right := 0
	for right < len(s) {
		in := s[right]
		right++

		window = append(window, in)
		for len(window) >= 10 {
			str := string(window)
			res[str]++
			if res[str] == 2 {
				ans = append(ans, str)
			}

			window = window[1:]
		}
	}
	return ans
}
