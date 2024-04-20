// @Author: Ciusyan 4/20/24

package cycle_6_4_17_4_21

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// 默认第一个字符串是最长的公共前缀
	comPrefix := strs[0]

	for cur := 1; cur < len(strs); cur++ {
		// 和当前字符挨个比较，看看新的公共前缀
		curStr := strs[cur]
		curL := min(len(comPrefix), len(curStr))
		// 看看能推到多远
		end := 0
		for i := 0; i < curL; i++ {
			if curStr[i] != comPrefix[i] {
				break
			}
			end++
		}
		if end == 0 {
			// 说明没有被推高，相当于 curStr 和 之前的 comPrefix 没有公共前缀，直接返回
			return ""
		}
		// 否则将推到的结果保存下来，去计算下一个字符串
		comPrefix = curStr[:end]
	}

	return comPrefix
}
