// @Author: Ciusyan 2023/7/27

package str

// https://leetcode.cn/problems/reverse-words-in-a-string/

func reverseWords(s string) string {
	if s == "" {
		return ""
	}

	// 方便处理字符串
	chars := []rune(s)

	// 1、去掉多余的空格
	l := removeSpace(chars)
	if l < 0 {
		return ""
	}

	// 2、整个逆序
	reverse(chars, 0, l)

	// 3、逆序每一个单词
	begin := -1
	for i := 0; i < l; i++ {
		if chars[i] == ' ' {
			reverse(chars, begin+1, i)
			begin = i
		}
	}

	// 还得对最后一个单词进行逆序
	reverse(chars, begin+1, l)

	return string(chars[:l])
}

// 去除多余空格，在原有基础上移动，并且返回对应合理的索引
func removeSpace(chars []rune) int {
	i, cur := 0, 0

	// 搞一个哨兵，哨兵位置默认让他是空格
	space := true
	for i < len(chars) {
		if chars[i] != ' ' {
			// 说明是有效字符
			space = false
			chars[cur] = chars[i]
			cur++
		} else if !space {
			// 说明是合法空格
			// 如果再遇到，那就代表是非法空格
			space = true
			chars[cur] = ' '
			cur++
		}

		i++
	}

	if space {
		// 说明最后是空格
		cur--
	}

	return cur
}

// 对 [begin, end) 的单词进行逆序
func reverse(chars []rune, begin, end int) {
	end--

	for begin < end {
		// 交换两个位置的字符
		chars[begin], chars[end] = chars[end], chars[begin]
		// 然后往中间靠
		begin++
		end--
	}
}
