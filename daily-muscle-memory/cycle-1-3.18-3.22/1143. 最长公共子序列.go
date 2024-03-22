// @Author: Ciusyan 3/22/24

package cycle_1_3_18_3_22

func longestCommonSubsequence(text1 string, text2 string) int {

	var (
		chars1 = []byte(text1)
		chars2 = []byte(text2)

		l1 = len(chars1)
		l2 = len(chars2)
	)

	// 准备 dp 缓存，dp[i1][i2] 代表：
	// chars1[0...i1) 和 chars[0...i2) 的最长公共子序列长度
	dp := make([][]int, l1+1)
	for i := range dp {
		dp[i] = make([]int, l2+1)
	}
	// dp[row][0] = dp[0][col] = 0
	for i1 := 1; i1 <= l1; i1++ {
		for i2 := 1; i2 <= l2; i2++ {
			// 四种情况取最大值， (i1, i2) (!i1, !i2) (i1, !i2) (!i1, i2)
			// 但是如果有 (i1, i2)，其余四种情况都一定没有它大
			// 如果没有 (i1, i2)， (!i1, !i2) (i1, !i2)，也一定比 (!i1, !i2) 大
			if chars1[i1-1] == chars2[i2-1] {
				// 说明有 (i1, i2)
				dp[i1][i2] = 1 + dp[i1-1][i2-1]
				continue
			}
			// 否则说明是 (!i1, !i2) (i1, !i2) 的最大值
			dp[i1][i2] = max(dp[i1-1][i2], dp[i1][i2-1])
		}
	}

	// 代表 chars1[0...l1) 和 chars[0...l2) 的最长公共子序列长度
	return dp[l1][l2]
}

func longestCommonSubsequence2(text1 string, text2 string) int {

	// chars1[0...i1) 和 chars2[0...i2 的最长公共子序列长度
	var process func(chars1, chars2 []byte, i1, i2 int) int
	process = func(chars1, chars2 []byte, i1, i2 int) int {
		if i1 == 0 || i2 == 0 {
			// 说明有一个字符串为 ""
			return 0
		}
		// 来到这里，至少证明两个序列都是有值的，有四种情况：
		// 选 i1 不选 i2
		p1 := process(chars1, chars2, i1, i2-1)
		// 选 i2 不选 i1
		p2 := process(chars1, chars2, i1-1, i2)
		// 两个都不选，但是这个情况其实不用考虑，因为最终是 p1、p2、p3、p4 取最大值，这个一定比不过上面两种情况
		// p3 := process(chars1, chars2, i1-1, i2-1)
		// 两个都选
		p4 := 0
		if chars1[i1-1] == chars2[i2-1] {
			// 能有第四种情况，说明最后一个字符相等了
			p4 = 1 + process(chars1, chars2, i1-1, i2-1)
		}

		return max(p1, p2, p4)
	}

	chars1 := []byte(text1)
	chars2 := []byte(text2)
	return process(chars1, chars2, len(chars1), len(chars2))
}
