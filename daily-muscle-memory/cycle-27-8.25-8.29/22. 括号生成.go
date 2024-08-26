// @Author: Ciusyan 2024/8/26

package cycle_27_8_25_8_29

// https://leetcode.cn/problems/generate-parentheses/description/

/*
*
思路重复：
使用 DFS 的方式即可，总共有 n 对括号，那么一个合法的字符串就是 2n 的长度。
那么我们从第 0 层开始，去收集括号，直至到达 2n 的长度后，就将其轨迹放入结果中，
那么如何收集呢？
只要有左括号，就可以进行收集，记录轨迹后去下一层搜索，并且将做括号余量减一
对于右括号，则需要比做括号多，才能够选择，否则选了也是无效括号，
循环往复，直至全部收集完毕，
*/

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
