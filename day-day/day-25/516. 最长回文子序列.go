// @Author: Ciusyan 11/22/23

package day_25

import "slices"

// https://leetcode.cn/problems/longest-palindromic-subsequence/description/

// 转换成求解最长公共子序列问题，比如：s = "bbbab"，s` = "babbb"，两个字符串的最大公共子序列是 bbbb，长度为 4
func longestPalindromeSubseq(s string) int {
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
