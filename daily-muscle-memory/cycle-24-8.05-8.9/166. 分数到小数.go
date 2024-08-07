// @Author: Ciusyan 2024/8/7

package cycle_24_8_05_8_9

import (
	"strconv"
	"strings"
)

// https://leetcode.cn/problems/fraction-to-recurring-decimal/description/

/**
思路重复：
1. 先除整数部分
如果没有余数说明能整除，否则需要处理小数部分，
2. 小数也是每次用余数 *10 再除被除数，但是我们每次都需要将余数给记下。
一旦出现过之前有过的余数，或者除数为零了，就可以终止了。
如果是前者，就说明有循环节，按照要求构建循环节即可。如果是后者，那就说明小数点后也能整除了。
*/

func fractionToDecimal(numerator int, denominator int) string {
	if denominator == 0 {
		return ""
	}

	if numerator == 0 {
		return "0"
	}

	sb := strings.Builder{}

	// 先处理符号
	if (numerator < 0) != (denominator < 0) {
		// 说明有一个是负号
		sb.WriteString("-")
		if numerator < 0 {
			numerator = -numerator
		} else {
			denominator = -denominator
		}
	}

	// 然后处理整数部分
	sb.WriteString(strconv.Itoa(numerator / denominator))
	numerator %= denominator
	if numerator == 0 {
		// 说明能整除
		return sb.String()
	}

	// 说明有小数
	sb.WriteString(".")

	// 但是得记录余数出现的次数，如果出现过相同的余数，那么就说明有循环节
	pre := make(map[int]int, 10) // pre[3] = 2 代表，3 之前出现过了，之后要添加括号，从字符串长度为 2 的位置开始加
	pre[numerator] = sb.Len()

	for numerator != 0 {
		// 借位再除
		numerator *= 10
		sb.WriteString(strconv.Itoa(numerator / denominator))
		numerator %= denominator

		// 检查之前是否出现过这个余数了
		preIdx, ok := pre[numerator]
		if !ok {
			// 说明之前没出现过，这下出现过了
			pre[numerator] = sb.Len()
		} else {
			// 出现过了，直接添加循环节即可
			tmp := []byte(sb.String())
			res := append([]byte{}, tmp[:preIdx]...)
			res = append(res, '(')
			res = append(res, tmp[preIdx:]...)
			res = append(res, ')')
			return string(res)
		}
	}

	return sb.String()
}
