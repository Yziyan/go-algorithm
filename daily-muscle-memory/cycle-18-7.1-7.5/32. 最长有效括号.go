// @Author: Ciusyan 2024/7/3

package cycle_18_7_1_7_5

// https://leetcode.cn/problems/longest-valid-parentheses/

func longestValidParentheses(s string) int {
	n := len(s)
	// dp[i] 代表，s[i] 结尾的字符串的最长有效括号长度
	dp := make([]int, n)

	res := 0
	for i := 1; i < n; i++ {
		if s[i] == '(' {
			// 以 s[i] 结尾，肯定凑不出有效括号了
			continue
		}
		// 有可能能凑出有效括号，先找到与之配对的 （
		pre := i - dp[i-1] - 1 // 当前位置 - 前一个位置有效括号的长度 - 1，就是对应的左括号
		if pre >= 0 && s[pre] == '(' {
			// 说明真的是左括号，长度至少为 2 了，但是还得加上 i-1 位置的有效括号长度
			dp[i] = 2 + dp[i-1]
			// 并且有可能 pre 前面还有有效括号，也要将其加上，比如 () | (())，如果 pre=2，它的前面还有一对有效括号
			if pre > 0 {
				dp[i] += dp[pre-1]
			}
		}
		// 每次都看看，能否推大结果
		if dp[i] > res {
			res = dp[i]
		}
	}

	return res
}
