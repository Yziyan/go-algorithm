// @Author: Ciusyan 1/20/24

package day_31

// https://leetcode.cn/problems/plus-one/

func plusOne(digits []int) []int {
	if digits == nil || len(digits) == 0 {
		return digits
	}
	l := len(digits)
	for i := l - 1; i >= 0; i-- {
		if digits[i] != 9 {
			// 说明这一位相加后，不会产生进位，加完直接返回即可
			digits[i]++
			return digits
		} else {
			// 要不然说明进位了
			digits[i] = 0
		}
	}

	// 能来到这里，说明所有数字都是 9，需要增加一位
	digits = make([]int, l+1)
	// 只有第一位是 1，比如 9999 + 1 -> 10000
	digits[0] = 1

	return digits
}
