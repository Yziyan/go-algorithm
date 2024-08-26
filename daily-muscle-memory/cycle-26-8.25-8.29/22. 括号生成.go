// @Author: Ciusyan 2024/8/26

package cycle_26_8_25_8_29

// https://leetcode.cn/problems/generate-parentheses/description/

func generateParenthesis(n int) []string {

	var dfs func(n int, level int, lL, rL int, trace []byte, res *[]string)
	dfs = func(n int, level int, lL, rL int, trace []byte, res *[]string) {
		if level == n {
			*res = append(*res, string(trace))
			return
		}

		if lL > 0 {
			// 左括号是有就能放
			trace[level] = '('
			dfs(n, level+1, lL-1, rL, trace, res)
		}

		// 右括号是必须要比左括号多才能使用
		if rL > lL {
			trace[level] = ')'
			dfs(n, level+1, lL, rL-1, trace, res)
		}
	}

	res := make([]string, 0, n)
	trace := make([]byte, n*2)
	dfs(n*2, 0, n, n, trace, &res)

	return res
}
