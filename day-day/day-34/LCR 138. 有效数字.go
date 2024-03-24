// @Author: Ciusyan 3/24/24

package day_34

import "strings"

func validNumber(s string) bool {
	// 用于去掉首尾空格
	l, r := 0, len(s)-1
	for l <= r && s[l] == ' ' {
		// 说明有前导空格
		l++
	}

	for r >= l && s[r] == ' ' {
		// 说明有后置空格
		r--
	}

	// 到达这里，说明前导和后置空格都去掉了，
	if l > r {
		// 到达这里，说明越界了
		return false
	}

	// 现在一定不存在合法的空格了
	s = s[l : r+1]
	// 然后看看是否是科学计数法表示的
	eCnt := strings.Count(s, "e") + strings.Count(s, "E")
	if eCnt > 1 {
		return false
	}
	if eCnt == 0 {
		// 说明不是科学计数法表示的，检查是否是小数或者整数
		return isInteger(s) || isDecimal(s)
	}

	// 来到这里，说明是科学计数法表示的，先找出 e 所在位置
	eIdx := max(strings.Index(s, "e"), strings.Index(s, "E"))

	// 科学计数法前的部分
	prefix := s[:eIdx]
	// 科学计数法后的部分
	postfix := s[eIdx+1:]

	// 前置部分可以允许是小数或者是整数，但是后置部分一定要求是整数
	return (isInteger(prefix) || isDecimal(prefix)) && isInteger(postfix)
}

// 是否是整数
func isInteger(s string) bool {
	if len(s) == 0 {
		return false
	}

	if s[0] == '-' || s[0] == '+' {
		// 去掉这个位置
		s = s[1:]
	}

	if len(s) == 0 {
		// 说明只有一个符号
		return false
	}

	for _, c := range s {
		if c < '0' || c > '9' {
			// 说明不是纯整数
			return false
		}
	}

	return true
}

// 是否是小数
func isDecimal(s string) bool {
	if len(s) == 0 {
		return false
	}
	// 判断符号，有符号
	if s[0] == '-' || s[0] == '+' {
		// 去掉这个位置
		s = s[1:]
	}

	if len(s) == 0 {
		// 说明只有一个符号
		return false
	}

	// 找出点的位置
	dotIdx := strings.Index(s, ".")
	if dotIdx == -1 {
		// 说明不是小数
		return false
	}

	// 否则可以拆分成两部分
	prefix := s[:dotIdx]
	postfix := s[dotIdx+1:]

	if len(prefix) == 0 && len(postfix) == 0 {
		// 说明是 "."
		return false
	}

	// 要求两部分都合法
	for _, c := range prefix {
		if c < '0' || c > '9' {
			// 说明前面不合法
			return false
		}
	}

	for _, c := range postfix {
		if c < '0' || c > '9' {
			// 说明后面不合法
			return false
		}
	}

	return true
}
