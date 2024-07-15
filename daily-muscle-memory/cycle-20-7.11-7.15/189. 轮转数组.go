// @Author: Ciusyan 2024/7/15

package cycle_20_7_11_7_15

// https://leetcode.cn/problems/rotate-array/

func rotate(nums []int, k int) {
	if len(nums) < 2 || k < 1 {
		return
	}

	// 对 nums 中的 [l...r) 进行逆序
	revers := func(nums []int, l, r int) {
		r--
		for l < r {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}
	}

	n := len(nums)
	k %= n

	// 1.先对 [0 ... n-k) 逆序
	revers(nums, 0, n-k)
	// 2.再对 [n-k ... n) 逆序
	revers(nums, n-k, n)
	// 3.最后再对 [0 ... n) 逆序
	revers(nums, 0, n)
}
