// @Author: Ciusyan 4/26/24

package cycle_7_4_22_4_26

/**
思路重复：
要对括号做匹配，可以使用栈来完成，
遇到左括号，我们就将其加入栈里，遇到右括号，就将栈顶元素弹出来进行匹配。
能匹配上菜进行后续的操作，当把所有括号都匹配完成后，看看栈里还有没有元素，有的话就说明不能完全匹配。
*/

func isValid(s string) bool {

	// 先做一个映射，方便比较
	m := map[rune]rune{
		'(': ')', '[': ']', '{': '}',
	}

	// 模拟栈
	stack := make([]rune, len(s))
	si := -1

	for _, c := range s {
		mc, ok := m[c]
		if ok {
			// 说明是左括号，压栈
			si++
			stack[si] = mc
		} else {
			// 说明是右括号，取出栈顶看看是否匹配
			if si == -1 {
				return false
			}

			top := stack[si]
			si--

			if top != c {
				return false
			}
		}
	}

	return si == -1
}
