// @Author: Ciusyan 5/25/24

package cycle_13_5_22_5_26

import (
	"fmt"
	"strings"
)

// https://leetcode.cn/problems/count-and-say/

/**
思路重复：
首先得读懂题目，
就是将一串字符，对字符用：出现次数+字符 表示。
比如：str = "231134441111133399" -> 可以表示为："1213211334513329"
这个关系可以使用一个数列构成，即：
prs(1) = "1"
prs(2) = "11"
prs(3) = "21"
prs(4) = "1211"
...

即，prs(n) 的值，是对于 prs(n-1) 的描述，这是一个很关键的信息
那么我们要算 prs(n)，就可以将 prs(n-1) 求出来，然后再对 prs(n-1) 的结果进行描述即可。
*/


func countAndSay(n int) string {
	if n <= 0 {
		return ""
	}

	if n == 1 {
		return "1"
	}

	// countAndSay(n) 是对 countAndSay(n-1) 结果的描述，那么我们得先知道要描述啥
	target := countAndSay(n - 1)
	sb := strings.Builder{}
	times := 1 // 看看每个字母出现了几次，出现的次数默认是一次
	l := len(target)

	for i := 1; i < l; i++ {
		if target[i] == target[i-1] {
			// 说明和之前一个字符相同，又出现了一次
			times++
		} else {
			// 说明不相同，记录上一个字符
			sb.WriteString(fmt.Sprintf("%d%c", times, target[i-1]))
			// 然后重置次数，去看下一个字符
			times = 1
		}
	}

	// 最后一个字符还没有人记录呢，因为我们只要到不相等了，才记录上一个字符，
	// 最后都退出了
	sb.WriteString(fmt.Sprintf("%d%c", times, target[l-1]))

	return sb.String()
}
