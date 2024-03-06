// @Author: Ciusyan 3/6/24

package day_33

import "strconv"

// https://leetcode.cn/problems/evaluate-reverse-polish-notation/description/

func evalRPN(tokens []string) int {
	if tokens == nil || len(tokens) == 0 {
		return 0
	}

	stack := make([]int, len(tokens))
	si := -1 // 代表栈顶

	for _, token := range tokens {
		if token == "+" || token == "-" || token == "*" || token == "/" {
			// 弹出两个元素
			num1 := stack[si]
			si--
			num2 := stack[si]
			si--

			// 看看是什么操作
			res := 0
			switch token {
			case "+":
				res = num2 + num1
			case "-":
				res = num2 - num1
			case "*":
				res = num2 * num1
			case "/":
				res = num2 / num1
			}

			// 需要将结果新入栈
			si++
			stack[si] = res
		} else {
			// 说明是数字，压入栈里
			si++
			num, _ := strconv.Atoi(token)
			stack[si] = num
		}
	}

	return stack[si]
}
