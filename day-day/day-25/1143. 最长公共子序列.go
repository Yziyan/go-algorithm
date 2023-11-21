// @Author: Ciusyan 11/21/23

package day_25

// https://leetcode.cn/problems/longest-common-subsequence/description/

// 动态规划方法
func longestCommonSubsequence(text1 string, text2 string) int {
	if text1 == "" || text2 == "" {
		return 0
	}

	chars1, chars2 := []byte(text1), []byte(text2)
	m, n := len(chars1), len(chars2)

	// 准备 DP 缓存，大小为：m * n，
	// dp[m][n] 代表，chars1[0 ... m] 和 chars2[0 ... n] 的最大公共子序列长度
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	// 填充 dp

	// 说明是 chars1[0] 和 chars2[0] 字符
	if chars1[0] == chars2[0] {
		// 并且还相等，那么最长公共子序列长度为 1
		dp[0][0] = 1
	}

	// 说明有  chars1[0] 和 chars2[0 ... n] 字符
	for col := 1; col < n; col++ {
		if chars1[0] == chars2[col] {
			dp[0][col] = 1
		} else {
			// 说明有没有 chars2[col] 字符都一样
			dp[0][col] = dp[0][col-1]
		}
	}

	// 说明有 chars1[0 ... m] 和 chars2[0] 字符
	for row := 1; row < m; row++ {
		if chars1[row] == chars2[0] {
			dp[row][0] = 1
		} else {
			// 说明有没有 chars1[row] 都一样
			dp[row][0] = dp[row-1][0]
		}
	}

	// 说明有 chars1[0 ... m] 和 chars2[0 ... m] 字符
	for row := 1; row < m; row++ {
		for col := 1; col < n; col++ {
			// p1：无 m 无 n，但是这种肯定没有下面几种大
			// p1 := dp[row-1][col-1]
			// p2：有 m 无 n
			p2 := dp[row][col-1]
			// p3：无 m 有 n
			p3 := dp[row-1][col]
			// p4：有 m 有 n，需要 chars1[m] 和 chars2[n] 相等
			p4 := 0
			if chars1[row] == chars2[col] {
				// 说明至少有一个长度了，看看剩余的有多长
				p4 = 1 + dp[row-1][col-1]
			}

			// 几种情况的最大值
			dp[row][col] = max(p2, p3, p4)
		}
	}

	// 那么结果应该是
	return dp[m-1][n-1]
}

// 暴力递归方法
func longestCommonSubsequence1(text1 string, text2 string) int {
	if text1 == "" || text2 == "" {
		return 0
	}

	// 递归含义：chars1[0 ... m] 和 chars2[0 ... n] 这两个范围的最大公共子序列的长度
	var process func(chars1, chars2 []byte, m, n int) int
	process = func(chars1, chars2 []byte, m, n int) int {
		if m == 0 && n == 0 {
			// 说明有 chars1[0] 和 chars2[0]
			if chars1[0] == chars2[0] {
				// 相等就返回 1
				return 1
			}

			return 0
		}

		if m == 0 {
			// 说明有 chars1[0] 和 chars2[0 ... n]
			if chars1[0] == chars2[n] {
				// 说明和 chars2 的最后一个字符相等
				return 1
			}

			// 否则有没有 chars2[n]，都不影响结果
			return process(chars1, chars2, m, n-1)
		}

		if n == 0 {
			// 说明有 chars1[0 ... m] 和 chars2[0]
			if chars1[m] == chars2[0] {
				// 说明和 chars1 的最后一个字符相等
				return 1
			}

			// 否则有没有 chars1[m]，都不影响结果
			return process(chars1, chars2, m-1, n)
		}

		// 能来到这里，说明有 chars1[0 ... m] 和 chars2[0 ... n]
		// 讨论 m 和 n 的样本即可
		// p1: 不 m，不 n，但是 p1 完全可以忽略掉，因为之后是 p1、p2、p3、p4 取最大值，它肯定 pk 不过 p2、p3，更不用说还可能有 p4 的可能了
		// p1 := process(chars1, chars2, m-1, n-1)
		// p2: 可能 m，不 n
		p2 := process(chars1, chars2, m, n-1)
		// p3: 不 m，可能 n
		p3 := process(chars1, chars2, m-1, n)
		// p4: 要 m，要 n，那就必须要求 chars1[m] == chars2[n]
		p4 := 0
		if chars1[m] == chars2[n] {
			// 说明 chars1[m] 和 chars2[n] 就占一个字符了，剩下的还能多长，就看缩小范围后能有多长了
			p4 = process(chars1, chars2, m-1, n-1) + 1
		}

		return max(p2, p3, p4)
	}

	chars1, chars2 := []byte(text1), []byte(text2)
	m, n := len(chars1)-1, len(chars2)-1

	// 代表求解：chars1[0 ... m] 和 chars2[0 ... n] 范围的最大公共子序列长度
	return process(chars1, chars2, m, n)
}
