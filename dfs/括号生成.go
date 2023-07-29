// @Author: Ciusyan 2023/7/29

package dfs

// https://leetcode.cn/problems/generate-parentheses/

func generateParenthesis(n int) []string {
	if n <= 0 {
		return nil
	}

	track := make([]rune, n<<1)
	var res []string

	// 从第 0 层开始，收集 2n 个括号，左右都剩余 n 个
	khDfs(0, n<<1, n, n, track, &res)

	return res
}

func khDfs(level, n, lRemain, rRemain int, track []rune, res *[]string) {
	if level == n {
		// 说明可以记录一个结果了
		*res = append(*res, string(track))

		return
	}

	// 有两种可能
	//	1、可选左括号
	if lRemain > 0 {
		// 只要有左括号，就能选左括号
		track[level] = '('
		// 往下钻
		khDfs(level+1, n, lRemain-1, rRemain, track, res)
	}

	// 	2、可选右括号
	if rRemain > lRemain {
		// 有右括号，并且右括号一定要比左括号多，
		track[level] = ')'
		khDfs(level+1, n, lRemain, rRemain-1, track, res)
	}
}
