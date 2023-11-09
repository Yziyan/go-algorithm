// @Author: Ciusyan 11/9/23

package day_24

// AllSubsquences 比如给定 str = "123" -> ["", "1", "2", "3", "12", "13", "23", "123"]
func AllSubsquences(str string) []string {
	if str == "" {
		return nil
	}

	res := make([]string, 0, 1)

	// 将其转换成 chars 数组
	chars := []byte(str)
	track := make([]byte, 0, len(chars))
	// level 在第几层，chars 原始字符数组 track 收集的轨迹 results 搜集的结果
	var dfs func(level int, chars []byte, track []byte, results *[]string)
	dfs = func(level int, chars []byte, track []byte, results *[]string) {
		if level == len(chars) {
			// 说明要收集一个结果
			*results = append(*results, string(track))
			return
		}

		// 列举所有肯呢个
		// 1.不选 chars[level]
		dfs(level+1, chars, track, results)

		// 2.选 chars[level]
		dfs(level+1, chars, append(track, chars[level]), results)
	}

	// 从第 0 层开始调用，收集结果
	dfs(0, chars, track, &res)

	return res
}
