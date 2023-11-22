// @Author: Ciusyan 11/22/23

package day_25

import "slices"

// https://leetcode.cn/problems/longest-palindromic-subsequence/description/

// 暴力递归方法，范围尝试模型
func longestPalindromeSubseq(s string) int {
	if s == "" {
		return 0
	}

	// 返回 chars[l ... r] 这个范围的最大回文子序列长度
	var process func(chars []byte, l, r int) int
	process = func(chars []byte, l, r int) int {
		if l == r {
			// 说明只有一个字符，那么必然回文
			return 1
		}

		if r-l == 1 {
			// 说明有两个字符
			if chars[l] == chars[r] {
				// 这俩字符还相等
				return 2
			}

			return 1
		}

		// 来到这里，说明大于 2 个字符，一般范围模型就考虑 l 和 r 位置
		// p1：无 l 无 r，因为需要求这几种情况的最大值，p1 肯定比比过 p2 和 p3，更别谈还可能有 p4
		// p1 := process(chars, l+1, r-1)
		// p2：无 l 有 r
		p2 := process(chars, l+1, r)
		// p3：有 l 无 r
		p3 := process(chars, l, r-1)
		// p4：有 l 有 r，但是这就要求 chars[l] 和 chars[r] 要相等
		p4 := 0
		if chars[l] == chars[r] {
			// 说明相等，那么就已经占 2 个字符了，再看看中间能得到的最大回文子序列长度有多少
			p4 = 2 + process(chars, l+1, r-1)
		}

		// 取几种情况的最大值
		return max(p2, p3, p4)
	}

	chars := []byte(s)
	// 返回 chars[0, len-1] 这个范围的最大回文子序列长度
	return process(chars, 0, len(s)-1)
}

// 转换成求解最长公共子序列问题，比如：s = "bbbab"，s` = "babbb"，两个字符串的最大公共子序列是 bbbb，长度为 4
func longestPalindromeSubseq1(s string) int {
	if s == "" {
		return 0
	}

	// 对 s 逆序
	chars1 := []byte(s)
	// 不能影响原切片
	chars2 := slices.Clone(chars1)
	// 需要先对原字符串逆序
	slices.Reverse(chars2)

	m, n := len(chars1), len(chars2)
	// 准备缓存 m*n
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// 填充 dp
	if chars1[0] == chars2[0] {
		dp[0][0] = 1
	}

	for col := 1; col < n; col++ {
		// 说明有 chars1[0] 和 chars2[1 ... n]
		if chars1[0] == chars2[col] {
			dp[0][col] = 1
		} else {
			// 说明没有 chars2[col] 也一样
			dp[0][col] = dp[0][col-1]
		}
	}

	for row := 1; row < m; row++ {
		// 说明有 chars1[1 ... m] 和 chars2[0]
		if chars1[row] == chars2[0] {
			dp[row][0] = 1
		} else {
			// 说明没有 chars1[row] 也一样
			dp[row][0] = dp[row-1][0]
		}
	}

	for row := 1; row < m; row++ {
		for col := 1; col < n; col++ {
			// p1 := dp[row-1][col-1]
			p2 := dp[row-1][col]
			p3 := dp[row][col-1]
			p4 := 0
			if chars1[row] == chars2[col] {
				p4 = 1 + dp[row-1][col-1]
			}
			dp[row][col] = max(p2, p3, p4)
		}
	}

	return dp[m-1][n-1]
}
