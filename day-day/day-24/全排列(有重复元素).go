// @Author: Ciusyan 11/10/23

package day_24

// 全排列2，所给的字符串，有可能有重复的，比如 "112" -> [112, 121, 211]

func PrintAllPermutationsNoRepeat(str string) []string {
	res := make([]string, 0, 1)
	if str == "" {
		return res
	}

	isRepeat := func(chars []byte, level, willSwap int) bool {
		for i := level; i < willSwap; i++ {
			if chars[i] == chars[willSwap] {
				return true
			}
		}

		return false
	}

	var dfs func(level int, chars []byte, res *[]string)
	dfs = func(level int, chars []byte, res *[]string) {
		if level == len(chars) {
			*res = append(*res, string(chars))
			return
		}

		for i := level; i < len(chars); i++ {
			// 看看前面是否有已经使用过的字符了，有了就不要换了
			if isRepeat(chars, level, i) {
				continue
			}

			// 确定 chars[level] 位置的元素
			chars[level], chars[i] = chars[i], chars[level]
			dfs(level+1, chars, res)
			// 还原现场
			chars[level], chars[i] = chars[i], chars[level]
		}
	}

	dfs(0, []byte(str), &res)
	return res
}
