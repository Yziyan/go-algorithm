// @Author: Ciusyan 11/16/23

package day_25

/**
规定 1和A 对应、2和B 对应、3和C 对应 ... 26和Z 对应
那么一个数字字符串比如 "111” 就可以转化为:
"AAA"、"KA" 和 "AK"
给定一个只有数字字符组成的字符串 str，返回有多少种转化结果
*/

// 暴力递归方法，str 只含有数字，返回最多的方法数，
func convertStrLetter1(str string) int {
	if str == "" {
		return 0
	}

	// 从左往右的尝试模型，含义是：当前处于 cur 位置，从 [cur ...] 开始，有多少种转换方法
	var process func(chars []byte, cur int) int
	process = func(chars []byte, cur int) int {
		if cur == len(chars) {
			// 说明能到达末尾，前面的肯定转换完成了，至少有一种转换方法
			return 1
		}

		if chars[cur] == '0' {
			// 如果遇到了 ‘0‘ 字符，说明没法转换
			return 0
		}

		// 否则要么选择一个字符转换、要么选择两个字符转换
		// 代表选择了当前字符，从 [cur+1 ...] 开始，得到的转换数
		p1 := process(chars, cur+1)

		p2 := 0
		if cur+1 < len(chars) && (chars[cur]-'0')*10+(chars[cur+1]-'0') < 27 {
			// 说明能选两个字符进行转换，比如选 11 -> J
			// 代表选择了两个字符，从 [cur+2 ...] 开始，得到的转换数
			p2 = process(chars, cur+2)
		}

		// 最终返回的结果就是，两种转换方法得到的转换数
		return p1 + p2
	}

	chars := []byte(str)
	// 那么返回值就是：当前处于 0 位置，从 [0 ...] 开始，有多少种转换方法
	return process(chars, 0)
}
