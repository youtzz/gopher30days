package problem0020

func isValid(s string) bool {
	// 字符串长度不是偶数说明一定是错误的: '(){'
	if len(s)%2 != 0 {
		return false
	}

	stack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		in := s[i]
		if in == '(' || in == '{' || in == '[' {
			stack = append(stack, in)
		} else {
			// 防止栈空时取元素
			if len(stack) == 0 {
				return false
			}
			out := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if (in == ')' && out != '(') || (in == ']' && out != '[') || (in == '}' && out != '{') {
				return false
			}
		}
	}

	if len(stack) > 0 {
		return false
	} else {
		return true
	}
}
