// @Author: Ciusyan 2024/7/2

package cycle_18_7_1_7_5

// https://leetcode.cn/problems/longest-palindromic-substring/

func longestPalindrome(s string) string {
	n := len(s)
	// dp[i][j] 代表，s [i ... j] 是回文子串
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)

		// 代表 s [i ... i] 是回文子串，只有一个字符当然是回文子串
		dp[i][i] = true
	}

	// 回文子串的开始位置
	begin := 0
	// 最长回文子串的长度
	maxLen := 1

	for row := n - 2; row >= 0; row-- {
		// 从对角线开始
		for col := row; col < n; col++ {
			curL := col - row + 1
			ht := s[row] == s[col]
			if curL <= 2 {
				// 说明只有一个字符或者两个字符，dp[i][j] 是回文子串就看仅有的字符是否相等。
				dp[row][col] = ht
			} else {
				// 说明多余两个字符，得看两个情况是不是相等的，先首尾，然后再看去掉首尾的字符串是不是回文子串
				dp[row][col] = ht && dp[row+1][col-1]
			}

			if dp[row][col] && curL > maxLen {
				// 说明最长回文子串可以更新了
				begin = row
				maxLen = curL
			}
		}
	}

	return s[begin : begin+maxLen]
}
