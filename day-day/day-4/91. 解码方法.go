// @Author: Ciusyan 2023/9/3

package day_4

// https://leetcode.cn/problems/decode-ways/

func numDecodings(s string) int {
	if s == "" {
		return 0
	}
	chars := []byte(s)
	l := len(chars)

	// dp[i] 代表以 s[i] 字符结尾，有多少种解码方法
	dp := make([]int, l+1)
	// "" -> 肯定有 1 种解码方案
	dp[0] = 1

	// 挨个往下递推
	for i := 1; i <= l; i++ {
		if chars[i-1] != '0' {
			// 说明至少可以解码
			dp[i] = dp[i-1]
		}

		// 尝试用前面两个字符能否推出一种方法，但是需要合法
		if i < 2 || chars[i-2] == '0' {
			// 前面都没有俩字符，或者往前数俩字符是 '0' 打头的，肯定不可以解码
			continue
		}

		// 还有往前数俩字符肯定不能大于 26
		if (chars[i-2]-'0')*10+chars[i-1]-'0' <= 26 {
			// 说明还可以解码，再加上一种方案
			dp[i] += dp[i-2]
		}
	}

	// 以第 s[l] 个字符结尾，有的解码方案
	return dp[l]
}

// 使用 DFS 的方法，但若数字太大，可能会超出时间限制
func numDecodings1(s string) int {
	if s == "" {
		return 0
	}

	// 从第 0 层开始搜索
	return decodeDfs(0, []byte(s))
}

// 从 level 层开始搜索，返回能够组合的方案次数
func decodeDfs(level int, chars []byte) int {
	if level == len(chars) {
		// 都搜到了末尾了，说明肯定是一种方案了
		return 1
	}

	// 如果 level 层的字符是 0 开头的，那么肯定不可以映射
	if chars[level] == '0' {
		return 0
	}

	// 说明当前字符没问题，可以搜索，找出可能的情况：
	//	1、先尝试一个字符往下搜索
	ways := decodeDfs(level+1, chars)

	// 	2、尝试两个字符往下搜索，但是要保证往下搜索合法
	//		首先是不能越界
	// 		还要保证这两个字符的组合不能大于 26，否则映射不出来
	if level+1 == len(chars) || (chars[level]-'0')*10+(chars[level+1]-'0') > 26 {
		return ways
	}

	// 到这里可以尝试两个字符，将两种方案得到的总次数汇总返回即是总次数
	return ways + decodeDfs(level+2, chars)
}
