// @Author: Ciusyan 1/31/24

package day_33

// https://leetcode.cn/problems/valid-palindrome/

func isPalindrome(s string) bool {
	if s == "" {
		return true
	}

	chars := []byte(s)
	var (
		n = len(chars)
		l = 0
		r = n - 1
	)

	// 双指针相遇了说明遍历完了，最后一个字符不需要比较了
	for l < r {

		for l < r {
			lValid, rValid := isValid(chars, l, r)
			if lValid && rValid {
				// 说明两个字符都合法了，终止遍历
				break
			}

			if !lValid {
				// 说明 l 位置的字符不合法，去下一个
				l++
			}
			if !rValid {
				// 说明 r 位置的字符不合法，去下一个
				r--
			}
		}

		// 来到这里，要么说明 l 和 r 的字符都合法了，要么说明越界了
		if l >= r {
			// 说明越界了，相当于都是无效字符，都相等
			return true
		}

		if !equals(chars[l], chars[r]) {
			// 只要有一次不相等，就不用看了
			return false
		}

		l++
		r--
	}

	return true
}

func equals(lc, rc byte) bool {
	if isChar(lc) && isChar(rc) {
		// 说明两个都是字母
		if lc == rc {
			// 说明大小写一致，并且相等
			return true
		}

		// 来到这里，要么是大小写不一致，要么是不相等
		if lc < rc {
			// 大小写相差 32
			return rc-lc == 32
		}

		// 说明 lc > rc 大小写相差 32
		return lc-rc == 32
	}

	// 来到这里，说明要么都是数字，要么至少有一个不是字母，
	// 不管怎样，都需要相等，才说明相等
	return lc == rc
}

// 求 l 和 r 是否是合法字符（数字 or 字母）
func isValid(chars []byte, l, r int) (bool, bool) {
	lValid := isChar(chars[l]) || isNum(chars[l])
	rValid := isChar(chars[r]) || isNum(chars[r])

	return lValid, rValid
}

// 是否是字母
func isChar(char byte) bool {
	if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
		return true
	}

	return false
}

// 是否是数字
func isNum(char byte) bool {
	if char >= '0' && char <= '9' {
		return true
	}
	return false
}
