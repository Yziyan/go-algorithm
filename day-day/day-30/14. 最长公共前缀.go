// @Author: Ciusyan 1/10/24

package day_30

// https://leetcode.cn/problems/longest-common-prefix/description/

func longestCommonPrefix(strs []string) string {
	if strs == nil || len(strs) == 0 {
		return ""
	}

	l := len(strs)
	// 默认将第一个字符串作为最长公共前缀
	maxPrefix := []byte(strs[0])
	for i := 1; i < l; i++ {
		// 每次和当前字符比较下，对照出一个最长的前缀来
		curStr := []byte(strs[i])
		if len(maxPrefix) > len(curStr) {
			// 说明 curStr 较短，最长也只能是他，所以用它作为基准
			maxPrefix, curStr = curStr, maxPrefix
		}

		// 对照每一个字符，记录最远能到达的索引
		commonIdx := 0
		for j, c := range maxPrefix {
			if c != curStr[j] {
				// 只要发现有不一样的字符了，说明前缀就不一样了
				break
			}
			commonIdx++
		}

		// 在结束的时候，把最新的公共前缀赋值后再去下一个字符
		maxPrefix = maxPrefix[:commonIdx]
	}

	return string(maxPrefix)
}
