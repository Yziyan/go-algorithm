// @Author: Ciusyan 1/15/24

package day_30

import (
	"fmt"
	"strings"
)

// https://leetcode.cn/problems/count-and-say/

func countAndSay(n int) string {
	if n <= 0 {
		return ""
	}

	if n == 1 {
		return "1"
	}

	// 描述 n，要先知道 n-1 是什么
	lastStr := countAndSay(n - 1)
	// 然后遍历 n-1 输出的字符串，来构建出 n 的字符串
	var (
		sb    strings.Builder
		times = 1 // 用于记录每一个数字出现的次数，默认第一个字符出现了一次
		l     = len(lastStr)
	)

	for i := 1; i < l; i++ {
		if lastStr[i] == lastStr[i-1] {
			// 如果和上一个字符相等，加一下次数即可
			times++
		} else {
			// 说明和上一个字符不一样了，记录上一个字符的答案
			sb.WriteString(fmt.Sprintf("%d%c", times, lastStr[i-1]))
			// 在填完了上一个数字的结果，需要重新计算次数
			times = 1
		}
	}

	// 遍历结束后，最后一个字符没有被记录，因为我们在循环里面，每次都是记录上一个字符
	sb.WriteString(fmt.Sprintf("%d%c", times, lastStr[l-1]))

	return sb.String()
}
