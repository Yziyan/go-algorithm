// @Author: Ciusyan 1/10/24

package day_30

// https://leetcode.cn/problems/letter-combinations-of-a-phone-number/description/

func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}

	// 先将所有数字的字符串映射整出来
	digitsMap := [][]byte{
		{'a', 'b', 'c'},      // 2
		{'d', 'e', 'f'},      // 3
		{'g', 'h', 'i'},      // 4
		{'j', 'k', 'l'},      // 5
		{'m', 'n', 'o'},      // 6
		{'p', 'q', 'r', 's'}, // 7
		{'t', 'u', 'v'},      // 8
		{'w', 'x', 'y', 'z'}, // 9
	}

	// DFS 函数部分，
	var dfs func(digits string, level int, track []byte, res *[]string)
	dfs = func(digits string, level int, track []byte, res *[]string) {
		if level == len(digits) {
			// 说明到达最后一层了，收集结果
			*res = append(*res, string(track))
			return
		}

		// 找出所有可能的字符串，因为没有 0和1，所以映射偏移了两个字符
		allChars := digitsMap[digits[level]-'2']

		for _, c := range allChars {
			// 收集当前字符
			track[level] = c
			// 然后往下一层钻，去收集下一层的字符
			dfs(digits, level+1, track, res)
		}
	}

	l := len(digits)
	// 排列组合，DSP 的套路
	res := make([]string, 0, l) // 结果
	track := make([]byte, l)    // 轨迹

	// 从第 0 层开始往下搜索
	dfs(digits, 0, track, &res)

	return res
}
