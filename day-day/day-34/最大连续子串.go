// @Author: Ciusyan 3/12/24

package day_34

/*
在无序的数字字符串中，找出连续的有序的最长串。比如 str = "3197345767852"，这个字符串中，最长的连续有序串是："345" ，或者 "678" 也可以。
*/

func findLongestOrderedSubstring(s string) string {
	if s == "" {
		return ""
	}

	chars := []byte(s)

	var (
		l        = len(chars)
		startIdx = 0
		maxLen   = 1
		curLen   = 1
	)

	for cur := 1; cur < l; cur++ {
		if chars[cur]-chars[cur-1] == 1 {
			// 说明是连续的
			curLen++
		} else {
			// 说明都不符合条件，将当前最大长度还原成默认值
			curLen = 1
		}

		if curLen <= maxLen {
			// 说明没有超过最长连续子串，不更新
			continue
		}

		// 来到这里，说明要更新了
		maxLen = curLen
		startIdx = cur - maxLen + 1
	}

	return string(chars[startIdx : startIdx+maxLen])
}
