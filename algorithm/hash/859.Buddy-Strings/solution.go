package problem0859

func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	// ("aa","aa")=true,("ab","ab")=false
	if s == goal {
		// 使用hashset统计字频
		set := make(map[rune]*struct{})
		for _, v := range goal {
			set[v] = nil
		}
		return len(s) > len(set)
	}

	var a, b string
	for i := 0; i < len(s); i++ {
		if s[i] != goal[i] {
			a = string(s[i]) + a // 这样处理方便后续判断a==b
			b += string(goal[i])
		}
	}
	return len(a) == 2 && a == b
}
