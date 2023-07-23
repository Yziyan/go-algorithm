// @Author: Ciusyan 2023/7/23

package other

func LengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}

	runes := []rune(s)
	if len(runes) == 0 {
		return 0
	}

	m := make(map[rune]int)

	begin := -1
	max := 1

	for i := 0; i < len(runes); i++ {
		endIdx, ok := m[runes[i]]
		if !ok {
			endIdx = -1
		}

		if begin < endIdx {
			// 说明在 begin 之后出现过了
			begin = endIdx
		} else {
			max = mathMax(max, i-begin)
		}

		m[runes[i]] = i
	}

	return max
}

func mathMax(a, b int) int {
	if a > b {
		return a
	}

	return b
}
