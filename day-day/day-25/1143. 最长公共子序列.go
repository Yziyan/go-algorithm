// @Author: Ciusyan 11/21/23

package day_25

// https://leetcode.cn/problems/longest-common-subsequence/description/

func longestCommonSubsequence(text1 string, text2 string) int {
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
		// p1: 不 m，不 n
		p1 := process(chars1, chars2, m-1, n-1)
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

		return max(p1, p2, p3, p4)
	}

	chars1, chars2 := []byte(text1), []byte(text2)
	m, n := len(chars1)-1, len(chars2)-1

	// 代表求解：chars1[0 ... m] 和 chars2[0 ... n] 范围的最大公共子序列长度
	return process(chars1, chars2, m, n)
}
