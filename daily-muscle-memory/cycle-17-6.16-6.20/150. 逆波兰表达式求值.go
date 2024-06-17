// @Author: Ciusyan 6/18/24

package cycle_16_6_11_6_15

import "strconv"

// https://leetcode.cn/problems/evaluate-reverse-polish-notation/description/

func evalRPN(tokens []string) int {
	stack := make([]int, len(tokens))
	si := -1 // 代表栈顶元素

	for _, t := range tokens {
		if t == "+" || t == "-" || t == "*" || t == "/" {
			// 说明需要弹出两个栈顶出来运算
			nums1 := stack[si]
			si--
			nums2 := stack[si]

			// 看看是何种运算符号
			switch t {
			case "+":
				stack[si] = nums2 + nums1
			case "-":
				stack[si] = nums2 - nums1
			case "*":
				stack[si] = nums2 * nums1
			case "/":
				stack[si] = nums2 / nums1
			}

		} else {
			// 说明是数字，直接入栈即可
			si++
			num, _ := strconv.Atoi(t)
			stack[si] = num
		}
	}

	if si == -1 {
		return 0
	}
	// 最后栈顶的就是答案
	return stack[si]
}
