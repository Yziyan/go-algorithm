// @Author: Ciusyan 2024/7/4

package cycle_18_7_1_7_5

// https://leetcode.cn/problems/valid-palindrome/description/

/*
思路重复：
如何看是否是回文串呢？其实使用双指针，挨个比较就可以了，但是这个题有一些限制，
比如有一些其他干扰字符，大小写得忽略。
所以，我们可以准备双指针，l,r 分别指向首尾，
然后先排除干扰字符，如果不是字母和数字，就直接跳过。
直至两个指针的位置都是合法的字符。
但是跳过后，如果已经相遇了，就不用比较了，说明全都是合法字符了。并且已经都没得比的了。
然后比较这两个位置的字符是否相等即可。
如何判断相等呢？其实就是看是不是相等，如果排除大小写的话，那就是差值得要是 32
*/

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1

	for l < r {

		for l < r {
			// 先跳过无效字符
			lv, rv := isValid(s, l, r)
			if lv && rv {
				// 说明都是有效字符了，直接去比较两个字符是否相等
				break
			}

			// 来到这里，总有一个不是有效的，直接跳过去下一个位置比较
			if !lv {
				// 左边不是有效的，跳过
				l++
			}

			if !rv {
				// 右边不是有效的，跳过
				r--
			}
		}

		if l >= r {
			// 说明跳过了无效字符后，只有一个字符或者没字符了，就说明是回文串
			return true
		}

		// 来到这里，看看两个字符是否是相等的，
		if !equal(s[l], s[r]) {
			// 说明不相等
			return false
		}

		l++
		r--
	}

	return true
}

// 判断 s[l] 和 s[r] 是否是有效字符（数字、字母）
func isValid(s string, l, r int) (lv, rv bool) {
	lv = isChar(s[l]) || isNum(s[l])
	rv = isChar(s[r]) || isNum(s[r])
	return
}

func equal(sl, sr byte) bool {
	if isChar(sl) && isChar(sr) {
		// 都是字符，就检查大小写即可，相等后面统一判定
		if sl < sr {
			return sr-sl == 32
		}

		if sl > sr {
			return sl-sr == 32
		}
	}

	// 来到这里，看看是否相等即可
	return sl == sr
}

func isNum(c byte) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}

func isChar(c byte) bool {
	if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
		return true
	}
	return false
}
