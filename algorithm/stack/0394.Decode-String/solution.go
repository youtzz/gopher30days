package problem0394

import (
	"strconv"
	"strings"
)

// 我的解法
func decodeString(s string) string {
	stack := make([]string, 0, len(s))
	sb := strings.Builder{}
	for i := 0; i < len(s); i++ {
		// 遇到数字入栈
		if s[i] >= '0' && s[i] <= '9' {
			stack = append(stack, string(s[i]))
		} else {
			// 如果是括号开始解析
			if s[i] == '[' {
				// 计算应该循环的次数
				times, _ := strconv.Atoi(strings.Join(stack, ""))
				// 清除栈，利用栈解析第一个[]中的内容
				stack = append(stack[len(stack):], "[")
				var str string
				for len(stack) > 0 {
					i++
					code := string(s[i])
					switch code {
					case "[":
						// 遇到 [ 入栈
						stack = append(stack, "[")
						str += string(code)
					case "]":
						// 遇到 ] 出栈
						stack = stack[:len(stack)-1]
						if len(stack) > 0 {
							str += code
						}
					default:
						str += code
					}
				}
				// 保存将解析好的字符串
				sb.WriteString(strings.Repeat(str, times))
				// 保存还未解析的字符串
				sb.WriteString(s[i+1:])

				s = sb.String()
				sb.Reset()
				i = -1 // i重新从0开始遍历，防止2[3[x]]这类情况
			} else {
				// 保存已解析的字符
				sb.WriteByte(s[i])
			}
		}
	}
	return s
}

// 最佳解法
func decodeString_best(s string) string {
	stk := make([]string, 0, len(s))
	ptr := 0
	for ptr < len(s) {
		code := s[ptr]
		if '0' <= code && code <= '9' {
			digit := getDigit(s, &ptr)
			stk = append(stk, digit)
		} else if ('a' <= code && code <= 'z') || ('A' <= code && code <= 'Z') || code == '[' {
			ptr++
			stk = append(stk, string(code))
		} else {
			ptr++
			var sub []string
			for stk[len(stk)-1] != "[" {
				sub = append(sub, "")
				copy(sub[1:], sub[:len(sub)-1])
				sub[0] = stk[len(stk)-1]
				stk = stk[:len(stk)-1]
			}
			times, _ := strconv.Atoi(stk[len(stk)-2])
			stk = stk[:len(stk)-2]
			str := strings.Repeat(strings.Join(sub, ""), times)
			stk = append(stk, str)
		}
	}
	return strings.Join(stk, "")
}

func getDigit(s string, ptr *int) string {
	var num string
	for *ptr < len(s) && s[*ptr] >= '0' && s[*ptr] <= '9' {
		num += string(s[*ptr])
		*ptr++
	}
	return num
}
