// @Author: Ciusyan 2023/8/2

package str

import (
	"fmt"
	"strings"
)

// 滴滴二面考题：大数乘法

func solution(str1, str2 string) string {

	chars1 := []rune(str1)
	chars2 := []rune(str2)

	// 选一个短的为 chars2
	if len(chars1) < len(chars2) {
		chars1, chars2 = chars2, chars1
	}

	// 将 chars1 转换成数字
	nums1 := 0
	s := 1
	for i := len(chars1) - 1; i >= 0; i-- {
		n := int(chars1[i] - '0')
		nums1 += n * s
		s *= 10
	}

	// 用于记录 chars2 计算后的结果
	m := make(map[rune]int, len(chars2))
	s = 1
	for i := len(chars2) - 1; i >= 0; i-- {
		n := int(chars2[i] - '0')

		m[chars2[i]] = nums1 * n * s
		s *= 10
	}

	res := 0
	for _, v := range m {
		res += v
	}

	return fmt.Sprint(res)
}

// 大数乘法，利用分治的思想
//
//			ab
//		*	cd
//		---------
//	        bd
//		  ad
//	      bc
//		 ac
//	 ab * cd = bd + ad + bc + ac
func multiply(str1, str2 string) string {
	// 递归基，如果某个数是 0，那么不用计算了
	if str1 == "0" || str2 == "0" {
		return "0"
	}

	// 递归基，如果某个数只有一位，那么就不用分治了
	if len(str1) == 1 {
		return strMultiply(str2, str1[0])
	} else if len(str2) == 1 {
		return strMultiply(str1, str2[0])
	}

	// 将 str1 和 str2 的对齐
	str1, str2 = fillZero(str1, str2)
	// 先计算出长度，方便相加时错位
	l := len(str1)
	mid := l >> 1

	a := str1[:mid] // str1 的高位
	b := str1[mid:] // str1 的低位
	c := str2[:mid] // str2 的高位
	d := str2[mid:] // str2 的低位

	// 分治的计算出低位相乘、高位相乘、交叉相乘
	ac := multiply(a, c) // 高位相乘
	ad := multiply(a, d) // 交叉 1
	bc := multiply(b, c) // 交叉 2
	bd := multiply(b, d) // 低位相乘

	// 不能直接将其相加，需要竖着相加，但是要先补 0
	ac = shiftLeft(ac, l)
	ad = shiftLeft(ad, mid)
	bc = shiftLeft(bc, mid)

	// 去除前导 0
	return strings.Trim(strAdd(strAdd(ac, bd), strAdd(ad, bc)), "0")
}

// 将 str1 和 str2 右对齐，左边用 0 填充
func fillZero(str1, str2 string) (string, string) {
	l := len(str1) - len(str2)

	if l > 0 {
		// 说明要对 str2 补齐 l 个 0
		str2 = strings.Repeat("0", l) + str2
	} else if l < 0 {
		// 说明要对 str1 补齐 -l 个 0
		str1 = strings.Repeat("0", -l) + str1
	}

	return str1, str2
}

// 对 str 右移 n 位，相当于在 str 后面，补充 n 个 0
func shiftLeft(str string, n int) string {
	if str == "0" {
		return str
	}

	return str + strings.Repeat("0", n)
}

// 对 str1 和 str2 相加
func strAdd(str1, str2 string) string {
	// 进位
	carry := byte(0)
	res := ""
	l1 := len(str1) - 1
	l2 := len(str2) - 1

	for l1 >= 0 || l2 >= 0 || carry > 0 {
		v1 := byte(0)
		if l1 >= 0 {
			v1 = str1[l1] - '0'
			l1--
		}

		v2 := byte(0)
		if l2 >= 0 {
			v2 = str2[l2] - '0'
			l2--
		}

		sum := v1 + v2 + carry
		// 更新进位
		carry = sum / 10
		// 取个位数
		sum %= 10

		// 拼接到结果的前面
		res = string(sum+'0') + res
	}

	return res
}

// str * c，一个数，乘单个数字
func strMultiply(str string, c byte) string {
	// 进位
	carry := byte(0)
	// 将单个字符转换成数字
	n := c - '0'
	res := ""
	l := len(str) - 1
	for l >= 0 || carry > 0 {
		v := byte(0)
		if l >= 0 {
			v = str[l] - '0'
			l--
		}
		// 相加
		v = v*n + carry
		// 更新进位
		carry = v / 10
		// 只要个位数
		v %= 10

		// 将结果拼接在前面
		res = string(v+'0') + res
	}

	return res
}
