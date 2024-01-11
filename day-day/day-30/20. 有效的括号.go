// @Author: Ciusyan 1/11/24

package day_30

// https://leetcode.cn/problems/valid-parentheses/

func isValid(s string) bool {
	if s == "" {
		return true
	}
	l := len(s)
	// 使用数组模拟一个栈
	stack := make([]rune, l)
	si := -1 // 代表栈顶元素

	for _, c := range s {
		if c == ')' || c == '}' || c == ']' {
			// 遇到了右括号，需要比较下是否匹配
			if si == -1 || stack[si] != c {
				// 说明栈里没有括号可以匹配，或者与当前括号不匹配
				return false
			}
			// 记得移动栈顶索引
			si--

			continue
		}

		// 来到下面，说明遇到的肯定是左括号了
		// 肯定需要压栈
		si++
		// 我们压相同方向的括号，到时候直接比较是否相等就可以看看是否匹配了
		if c == '(' {
			stack[si] = ')'
		} else if c == '{' {
			stack[si] = '}'
		} else { // 说明是 '['
			stack[si] = ']'
		}
	}

	// 最后要求栈一定是空的，要不然说明没有匹配完
	return si == -1
}
