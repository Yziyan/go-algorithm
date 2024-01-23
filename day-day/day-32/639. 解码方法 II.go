// @Author: Ciusyan 1/23/24

package day_32

// https://leetcode.cn/problems/decode-ways-ii/

// 动态规划方法
func numDecodingsII1(s string) int {
	if s == "" {
		return 0
	}

	n := len(s)

	// 准备缓存 dp，dp(cur) 代表：解码 s[cur ...] 的方法数
	dp := make([]int64, n+1)

	// 递归基可知
	dp[n] = 1

	mod := int64(1000000000 + 7)

	// dp(cur) 依赖 dp(cur+1)，所以从后往前求解
	for cur := n - 1; cur >= 0; cur-- {
		curC := s[cur]

		if curC == '0' {
			// 不能 '0' 打头
			// dp[cur] = 0
			continue
		}

		if curC == '*' {
			// 说明以 * 打头
			// 单个字符有：1~9，9种可能
			p1 := 9 * dp[cur+1]

			// 尝试使用两个字符
			if cur+1 == n {
				dp[cur] = p1 % mod
				continue
			}

			// 说明有两个字符
			p2 := int64(0)
			cur1C := s[cur+1]
			if cur1C != '*' {
				// 说明是 *0 ~ *9
				if cur1C > '6' {
					// 说明是 *7 ~ *9
					p2 = dp[cur+2]
				} else {
					// 说明是 *0 ~ *6
					p2 = 2 * dp[cur+2]
				}
			} else {
				// 说明是 **
				p2 = 15 * dp[cur+2]
			}

			dp[cur] = (p1 + p2) % mod
			continue
		}

		// 来到这里，说明不是 * 打头
		// 那么肯定能用一个字符解码了
		p1 := dp[cur+1]

		// 尝试使用两个字符解码
		if cur+1 == n {
			dp[cur] = p1 % mod
			continue
		}
		// 说明有两个字符
		p2 := int64(0)
		cur1C := s[cur+1]

		if cur1C == '*' {
			// 说明是 x*，但是只有 1{1~9} 和 2{1~6} 才有效
			if curC == '1' {
				p2 = 9 * dp[cur+2]
			} else if curC == '2' {
				p2 = 6 * dp[cur+2]
			}
		} else {
			// 说明是 xx
			if ((curC-'0')*10 + cur1C - '0') < 27 {
				// 那么这两个字符必须要合理
				p2 = dp[cur+2]
			}
		}

		dp[cur] = (p1 + p2) % mod
	}

	// 解码 s[0 ...] 的方法数
	return int(dp[0])
}

// 暴力递归方法
func numDecodingsII(s string) int {
	if s == "" {
		return 0
	}

	// 解码 s[cur ...] 所拥有的解码数量
	var process func(s string, cur int) int
	process = func(s string, cur int) int {
		if cur == len(s) {
			// 说明成功有一种解码数量了
			return 1
		}

		curC := s[cur]
		if curC == '0' {
			// 不能独自面对 '0'
			return 0
		}

		if curC == '*' {
			// 是 * 打头，
			// 所以如果选一个字符，有 9 种可能
			p1 := 9 * process(s, cur+1)

			// 尝试使用两个字符解码，那么只能是 1x 和 2x 才可能有效，
			if cur+1 == len(s) {
				// 说明没有两个字符
				return p1
			}

			p2 := 0
			cur1C := s[cur+1]
			if cur1C == '*' {
				// 说明是 **
				// 11 ~ 19: 9
				// 21 ~ 26: 6
				p2 = 15 * process(s, cur+2)
			} else {
				// 说明是 *{0~9}
				// 10 20 | 11 21 | 12 22 | 13 23 | 14 24 | 15 25 | 16 26 | 17 | 18 | 19
				// 所以
				if cur1C > '6' {
					p2 = process(s, cur+2)
				} else {
					// 说明是 0~6
					p2 = 2 * process(s, cur+2)
				}
			}

			return p1 + p2
		}

		// 来到这里，说明不是 '*' 打头
		// 那么肯定能够使用单独字符转了
		p1 := process(s, cur+1)

		// 第二种情况是：使用两个字符转
		if cur+1 == len(s) {
			// 说明没有两个字符了
			return p1
		}

		p2 := 0
		// 有两个字符，可能是 1{0~9} or 1*
		cur1C := s[cur+1]
		if cur1C != '*' {
			// 说明是 cur+1 正常数字
			if ((curC-'0')*10 + cur1C - '0') < 27 {
				// 说明两位数合法
				p2 = process(s, cur+2)
			}
		} else {
			if curC < '3' {
				// 只能有 1* 和 2*，3* 肯定不合法了
				if curC == '1' {
					// 说明是 1*，* 能代表 (1 ~ 9)
					p2 = 9 * process(s, cur+2)
				} else {
					// 说明是 2*，* 能代表 (1~6)
					p2 = 6 * process(s, cur+2)
				}
			}
		}

		return p1 + p2
	}

	// 解码 s[0 ...] 拥有的解码数量
	return process(s, 0) % (1000000000 + 7)
}
