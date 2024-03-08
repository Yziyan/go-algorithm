// @Author: Ciusyan 3/8/24

package day_34

import (
	"strconv"
	"strings"
)

/**
题意：
● 给定一个有序数组 num：[0, 2, 3, 10, 50, 78]
● 给定一个 lower：-3
● 给定一个 upper：99
● 需要返回相隔为 1 的编辑距离：["-3->-1", "1", "4->9", "11->49", "51->77", "79->99"]
*/

func findMissingRanges(nums []int, lower, upper int) []string {
	res := make([]string, 0, len(nums))

	// 遍历每一个数字
	for _, num := range nums {
		if num > lower {
			// 说明 num 前的数，都有距离
			res = append(res, miss(lower, num-1))
		}

		if num == upper {
			// 说明已经到达 upper 了，不需要看其他数了
			return res
		}

		lower = num + 1
	}

	// 在遍历完之后，还有可能最后一段还么有达到高度
	if lower <= upper {
		// 说明还需要增加最后一段
		res = append(res, miss(lower, upper))
	}

	return res
}

func miss(lower, upper int) string {
	var sb strings.Builder
	// 添加左边界
	sb.WriteString(strconv.Itoa(lower))
	if lower == upper {
		// 说明刚好距离就是 1，直接返回即可
		return sb.String()
	}
	// 否则说明需要添加距离符号："->"
	sb.WriteString("->")
	// 添加右边界
	sb.WriteString(strconv.Itoa(upper))

	return sb.String()
}
