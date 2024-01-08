// @Author: Ciusyan 1/8/24

package day_30

import "math"

// https://leetcode.cn/problems/string-to-integer-atoi/

func myAtoi(s string) int {
	// 转换成合法的字符串，第二个返回值代表能否转换
	chars := []rune(s)
	// 开始转换的索引
	idx := 0
	n := len(chars)

	// 去掉前导 ' '
	for idx < n && chars[idx] == ' ' {
		idx++
	}

	if idx == n {
		// 说明去掉前导空格后，都没有字符了
		return 0
	}

	// 到达这里，说明可以转换了，那么将其挨个位置转换
	var (
		res = 0
		// 是否是负数
		neg = false

		// 下面两个参数，用于判断是否溢出
		minDiv10 = math.MinInt32 / 10
		minMod10 = math.MinInt32 % 10
	)

	if chars[idx] == '+' || chars[idx] == '-' {
		if chars[idx] == '-' {
			// 标记为负数
			neg = true
		}
		// 说明带符号，往后一位开始
		idx++
	}

	for idx < n {
		c := chars[idx]
		if c < '0' || c > '9' {
			// 说明遇到了不是数字的字符1，就没必要转换了
			break
		}
		// 先将当前字符转换成负数，因为负数的范围更大
		curNum := int('0' - c)

		// 检查是否会溢出
		if res < minDiv10 || (res == minDiv10 && curNum < minMod10) {
			// 说明当前的 res：
			// 1.比系统最小除10还小，那么稍后乘10，肯定会小于系统最小值，即溢出
			// 2.等于系统最小值除10，那么还得看加上当前数字后，会比系统最小值还小，即溢出
			if neg {
				return math.MinInt32
			}
			return math.MaxInt32
		}

		// 来到这里说明不会溢出，那么加上当前数字
		res = res*10 + curNum
		idx++
	}

	if !neg && res == math.MinInt32 {
		// 说明转换的数字是正数，但是转换出来后是系统最小值，那么转换成整数就会溢出
		return math.MaxInt32
	}

	// 否则返回结果
	if neg {
		return res
	}
	return -res
}
