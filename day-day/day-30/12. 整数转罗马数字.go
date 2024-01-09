// @Author: Ciusyan 1/9/24

package day_30

import "strings"

// https://leetcode.cn/problems/integer-to-roman/description/

func intToRoman(num int) string {
	// 准备一张映射表
	romanNumMap := [][]string{
		{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}, // 个位
		{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}, // 十位
		{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}, // 百位
		{"", "M", "MM", "MMM"}, // 千位，因为这题最大是三千，所以后面的不用写了
	}

	// 拼接结果
	sb := strings.Builder{}
	// 先拼千位，取出千位数字
	sb.WriteString(romanNumMap[3][num/1000%10])
	// 拼百位，取出百位数字
	sb.WriteString(romanNumMap[2][num/100%10])
	// 拼十位，取出十位数字
	sb.WriteString(romanNumMap[1][num/10%10])
	// 最后拼个位，取出个位数字
	sb.WriteString(romanNumMap[0][num/1%10])

	return sb.String()
}
