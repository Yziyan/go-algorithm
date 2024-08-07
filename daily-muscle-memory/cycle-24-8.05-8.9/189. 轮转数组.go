// @Author: Ciusyan 2024/7/15

package cycle_24_8_05_8_9

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
