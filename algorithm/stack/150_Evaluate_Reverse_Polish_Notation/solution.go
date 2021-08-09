package main

import (
	"strconv"
)

// first solution
func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	for _, v := range tokens {
		if v == "+" || v == "-" || v == "*" || v == "/" {
			num1, num2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]

			var rst int
			switch v {
			case "+":
				rst = num1 + num2
			case "*":
				rst = num1 * num2
			case "-":
				rst = num1 - num2
			case "/":
				rst = num1 / num2
			}

			stack = append(stack, rst)
		} else {
			num, _ := strconv.Atoi(v)
			stack = append(stack, num)
		}
	}

	return stack[0]
}
