// @Author: Ciusyan 1/23/24

package day_32

// https://leetcode.cn/problems/decode-ways/

// 动态规划解法
func numDecodings(s string) int {
	if s == "" {
		return 0
	}

	n := len(s)

	// 准备缓存 dp，dp[cur] 的含义是：
	// 转换 s[cur ...] 所拥有的方法数
	dp := make([]int, n+1)

	// 根据递归基
	dp[n] = 1

	// 因为 dp(cur) 依赖 dp[cur+1]，所以从后往前求解
	for cur := n - 1; cur >= 0; cur-- {
		curC := s[cur]

		if curC == '0' {
			// 说明当前独立面对了 '0' 字符
			// dp[cur] = 0
			continue
		}

		// 两种可能
		// 1.选择一个字符转
		p1 := dp[cur+1]

		// 2.选择两个字符转
		p2 := 0
		// 转换成数字的时候，记得这里谁在前谁在后
		if cur+1 < n && ((curC-'0')*10+s[cur+1]-'0') < 27 {
			// 说明有两个字符可以转，并且转换也合法
			p2 = dp[cur+2]
		}

		dp[cur] = p1 + p2
	}

	// 代表从 s[0 ...] 拥有的转换数
	return dp[0]
}

// 暴力递归解法
func numDecodings1(s string) int {
	if s == "" {
		return 0
	}

	// 憋一个暴力递归，代表：
	// 要去解析 s[cur ...]，有多少种方法数
	var process func(s string, cur int) int
	process = func(s string, cur int) int {
		if cur == len(s) {
			// 说明成功有一种解析方式了
			return 1
		}

		// 先取出当前字符
		curC := s[cur]

		if curC == '0' {
			// 说明独自面对了 '0' 字符，转不了一点
			return 0
		}

		// 那么有两种转换方式
		// 1.选当前一个字符
		p1 := process(s, cur+1)

		// 2.选两个字符转
		if cur+1 == len(s) {
			// 说明没有两个字符了
			return p1
		}

		// 将两个字符转换成数字
		num := (curC-'0')*10 + s[cur+1] - '0'
		if num > 26 {
			// 说明不能转
			return p1
		}
		p2 := process(s, cur+2)

		// 两种可能的结果返回即可
		return p1 + p2
	}

	// 从 s[0 ...] 开始去解析，拥有的方法数
	return process(s, 0)
}
