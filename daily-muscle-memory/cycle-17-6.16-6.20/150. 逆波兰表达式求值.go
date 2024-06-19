// @Author: Ciusyan 6/18/24

package cycle_16_6_11_6_15

import "strconv"

// https://leetcode.cn/problems/evaluate-reverse-polish-notation/description/

/**
思路重复：
如果读懂题目后，其实也不是特别难，有点那种纸老虎的感觉。
即使用栈来操作即可，核心思路：
1. 当遇到数字，直接压入栈中，等待弹出计算。
2. 当遇到操作符，直接弹出两个栈顶元素，使用符号做对应的计算。并且将得到的结果压入栈中
然后最后当所有字符都计算完成后，栈顶元素就是最终的结果。
*/

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
