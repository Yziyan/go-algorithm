// @Author: Ciusyan 1/12/24

package day_30

// https://leetcode.cn/problems/generate-parentheses/description/

func generateParenthesis(n int) []string {
	if n <= 0 {
		return nil
	}

	// level 第几层 l 轨迹长度 leftRemain、rightRemain 左括号和右括号剩余的数量
	// track 轨迹 res 收集结果
	var dfs func(level int, l int, leftRemain, rightRemain int, track []rune, res *[]string)
	dfs = func(level int, l int, leftRemain, rightRemain int, track []rune, res *[]string) {
		if level == l {
			// 说明到达了最后一层，收集结果
			*res = append(*res, string(track))
			return
		}

		// 来到这里，要么选左括号，要么选右括号
		if leftRemain > 0 {
			// 只要有左括号，就能选
			track[level] = '('
			// 选择后去下一层搜索，剩余的左括号数量就得减一了
			dfs(level+1, l, leftRemain-1, rightRemain, track, res)
		}

		if rightRemain > leftRemain {
			// 右括号比左括号多才能选，要不然选了会是无效的组合
			track[level] = ')'
			// 选择后去下一层搜索，剩余的右括号数量就得减一了
			dfs(level+1, l, leftRemain, rightRemain-1, track, res)
		}
	}

	// 完整的轨迹肯定是将 n 对括号都装下，所以长度是 2n
	l := n << 1
	res := make([]string, 0, l)
	track := make([]rune, l)

	// 从第 0 层开始搜索，一开始左、右括号都剩余 n 个
	dfs(0, l, n, n, track, &res)

	return res
}
