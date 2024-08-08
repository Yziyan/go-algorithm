// @Author: Ciusyan 2024/7/15

package cycle_24_8_05_8_9

import (
	"strconv"
	"strings"
)

// https://leetcode.cn/problems/rotate-array/

/**
思路重复：
三个步骤：
1.先对 [0, n-k) 翻转
2.在对 [n-k, n) 翻转
3.最后对 [0, n) 翻转

但是 k 必须不能越界，也就是必须要翻转前，对数组长度取一下模
*/

func rotate(nums []int, k int) {
	if len(nums) < 2 && k < 1 {
		return
	}

	reverse := func(nums []int, l, r int) {
		r--
		for l < r {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}
	}

	n := len(nums)

	// 防止轮转的长度比数组长度还长
	k %= n

	// 对 [0 ... k) 逆序
	reverse(nums, 0, n-k)
	// 对 [k ... n) 逆序
	reverse(nums, n-k, n)
	// 对 [0 ... n) 逆序
	reverse(nums, 0, n)
}

/**
思路重复：
1. 先除整数部分
如果没有余数说明能整除，否则需要处理小数部分，
2. 小数也是每次用余数 *10 再除被除数，但是我们每次都需要将余数给记下。
一旦出现过之前有过的余数，或者除数为零了，就可以终止了。
如果是前者，就说明有循环节，按照要求构建循环节即可。如果是后者，那就说明小数点后也能整除了。
*/

func prs(t, b int) string {
	if t == 0 {
		return "0"
	}

	if b == 0 {
		return ""
	}

	sb := strings.Builder{}
	if (t < 0) != (b < 0) {
		sb.WriteString("-")
		if t < 0 {
			t = -t
		} else {
			b = -b
		}
	}

	sb.WriteString(strconv.Itoa(t / b))
	t %= b
	pre := make(map[int]int, 10)
	pre[t] = sb.Len()

	for t != 0 {
		t *= 10
		sb.WriteString(strconv.Itoa(t / b))
		t %= b

		if idx, ok := pre[t]; !ok {
			pre[t] = sb.Len()
		} else {
			tmp := []byte(sb.String())
			res := append([]byte{}, tmp[:idx]...)
			res = append(res, '(')
			res = append(res, tmp[idx:]...)
			res = append(res, ')')
			return string(res)
		}
	}

	return sb.String()
}
