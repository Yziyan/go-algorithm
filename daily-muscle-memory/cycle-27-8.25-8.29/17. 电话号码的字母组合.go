// @Author: Ciusyan 2024/8/27

package cycle_27_8_25_8_29

// https://leetcode.cn/problems/letter-combinations-of-a-phone-number/description/

/**
思路重复：
每一个数字，都有能够选择的字符，有多少种组合呢？使用 DFS 即可做了。
我们只需要将每一个数字对应的字符做一个映射即可，
那么我们从第一层开始去搜索，知道收到了 digits 这么长的轨迹，
到达每一层，轨迹都可以使用对应数字的字符。然后选择一个后，去到下一层即可。
*/

func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}

	dMap := [][]byte{
		{'a', 'b', 'c'},      // 2
		{'d', 'e', 'f'},      // 3
		{'g', 'h', 'i'},      // 4
		{'j', 'k', 'l'},      // 5
		{'m', 'n', 'o'},      // 6
		{'p', 'q', 'r', 's'}, // 7
		{'t', 'u', 'v'},      // 8
		{'w', 'x', 'y', 'z'}, // 9
	}

	var dfs func(level int, digits string, track []byte, res *[]string)
	dfs = func(level int, digits string, track []byte, res *[]string) {
		if level == len(digits) {
			*res = append(*res, string(track))
			return
		}

		// 有哪些选择？看看映射出了哪些字符
		strs := dMap[digits[level]-'2']

		for _, str := range strs {
			track[level] = str
			dfs(level+1, digits, track, res)
		}
	}

	res := make([]string, 0, 100)
	track := make([]byte, len(digits))
	dfs(0, digits, track, &res)

	return res
}
