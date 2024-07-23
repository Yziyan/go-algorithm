// @Author: Ciusyan 2024/7/18

package cycle_21_7_16_7_20

/**
思路重复：
有多少种解码方法呢？我们应该可以使用 dp 的方式求。
准备 dp = []int
dp[cur] 代表，使用 s[cur ...] 有多少种解码方法。
对于 cur 位置而言，只要 s[cur] 不是 0，就至少有一种解码方法，
即使用 dp[cur+1] 即可，但是如果要用两位来解码，那么就得看 s[cur ... cur+1] 不大于 26，
即可拥有一种方法，使用 dp[cur+2] 即可
*/

// https://leetcode.cn/problems/decode-ways/description/
func numDecodings(s string) int {
	n := len(s)
	// dp[cur] 代表，从 s[cur ...] 拥有的解码数量
	dp := make([]int, n+1)
	dp[n] = 1 // 代表 s[n ...] 的解码数量，相当于是 "" 的解码数量

	for cur := n - 1; cur >= 0; cur-- {
		curC := s[cur] - '0'
		if curC <= 0 {
			// 说明零打头，肯定不能解码了
			continue
		}
		// 既然当前字符能选，那么后续的有多少种，就看先前计算的即可
		dp[cur] = dp[cur+1]
		if cur+1 < n && curC*10+s[cur+1]-'0' < 27 {
			// 说明还能选两个数字的组合，那么选择两个数字，后续怎么的就看先前计算的了
			dp[cur] += dp[cur+2]
		}
	}

	// 代表从 s[0 ...] 的解码数量
	return dp[0]
}

func numDecodings1(s string) int {
	var prs func(s string, cur int) int
	prs = func(s string, cur int) int {
		if cur == len(s) {
			return 1
		}
		curC := s[cur] - '0'
		if curC <= 0 {
			return 0
		}

		res := prs(s, cur+1)
		if cur+1 < len(s) && curC*10+s[cur+1]-'0' < 27 {
			res += prs(s, cur+2)
		}
		return res
	}

	return prs(s, 0)
}
