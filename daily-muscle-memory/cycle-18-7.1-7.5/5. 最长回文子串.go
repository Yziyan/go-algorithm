// @Author: Ciusyan 2024/7/2

package cycle_18_7_1_7_5

// https://leetcode.cn/problems/longest-palindromic-substring/

/**
思路重复：
我们首先找到这个字符串，有多少回文子串。
然后我们对所有的回文子串的长度进行比较，找出最大的那个即可。
我们可以使用 dp 的方式，来进行判断，是否是回文子串。
即当 dp[row][col] 代表着 s[row ... col] 是否是回文子串。
那么 dp[row][row] 肯定是回文子串，因为 s[row] == s[row]，这就是初始条件
那么一般情况呢？当 col-row+1 <= 2 时，相当于最多只有两个端点，那么如果 s[row] == s[col]，即是回文子串。
否则首先看端点，然后看看 dp[row+1][col-1]，这个去掉端点后的是不是回文子串，如果是，才能说明整体是回文子串。

当计算出一个回文子串后，就看看能不能更新结果的长度？还有回文串结果的起始点。
*/
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
