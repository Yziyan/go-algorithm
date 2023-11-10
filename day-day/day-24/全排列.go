// @Author: Ciusyan 11/9/23

package day_24

// PrintAllPermutations 全排列，所有元素都需要选择， 123 -> 123、132、213、231、312、321
// 使用 DFS 来解答
func PrintAllPermutations(str string) []string {
	res := make([]string, 0, 1)
	if str == "" {
		return res
	}

	// chars[level] 是当前字符, track 用于暂存轨迹, used 用于记录哪些字符已经使用过了, res 用于收集结果
	var dfs func(level int, chars []byte, track []byte, used []bool, res *[]string)
	dfs = func(level int, chars []byte, track []byte, used []bool, res *[]string) {
		if level == len(chars) {
			// 说明到最后一个字符了，需要收集一个结果
			*res = append(*res, string(track))
			return
		}

		// 否则列举所有的可能，但是需要标识，哪些字符已经使用过了
		for i := 0; i < len(chars); i++ {
			if used[i] {
				// 代表使用过了
				continue
			}
			// 记录轨迹
			track[level] = chars[i]
			used[i] = true
			// 往下钻
			dfs(level+1, chars, track, used, res)
			// 但是需要还原现场
			used[i] = false
		}
	}

	// 转换成字符数组
	chars := []byte(str)
	// 从第一个字符开始，每次收集的结果暂存在 track 里面，将收集的结果都放置在 res 里面
	dfs(0, chars, make([]byte, len(chars)), make([]bool, len(chars)), &res)

	return res
}
