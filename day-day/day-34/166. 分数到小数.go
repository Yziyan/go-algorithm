// @Author: Ciusyan 3/8/24

package day_34

import (
	"strconv"
	"strings"
)

func fractionToDecimal(numerator int, denominator int) string {
	if denominator == 0 {
		return ""
	}

	if numerator == 0 {
		return "0"
	}

	var sb strings.Builder
	// 先处理符号
	if (numerator < 0) != (denominator < 0) {
		// 说明有一个是负数，添加符号
		sb.WriteByte('-')
		if numerator < 0 {
			numerator = -numerator
		}
		if denominator < 0 {
			denominator = -denominator
		}
	}

	// 先添加整数部分
	sb.WriteString(strconv.Itoa(numerator / denominator))

	numerator %= denominator

	if numerator == 0 {
		return sb.String()
	}
	// 来到这里，说明有小数产生，需要计算小数部分
	sb.WriteByte('.')
	// 用于标记前面是否已经计算过这个小数了
	// preNum[3] = 2 代表下一个余数，前面在 idx = 2 时，已经出现过了，出现了循环节
	preNum := make(map[int]int)
	preNum[numerator] = sb.Len()

	// 只要为 0 了，就说明没有循环小数
	for numerator != 0 {
		numerator *= 10
		sb.WriteString(strconv.Itoa(numerator / denominator))
		numerator %= denominator

		preIdx, ok := preNum[numerator]
		if !ok {
			// 说明之前没有计算过这个小数
			preNum[numerator] = sb.Len()
		} else {
			// 说明之前计算过这个数，出现了循环节
			tempRes := []byte(sb.String())
			res := append([]byte(nil), tempRes[:preIdx]...)
			res = append(res, '(')
			res = append(res, tempRes[preIdx:]...)
			res = append(res, ')')

			return string(res)
		}
	}

	return sb.String()
}
