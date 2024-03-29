// @Author: Ciusyan 3/29/24

package cycle_4_3_29_4_3

// https://leetcode.cn/problems/move-zeroes/

func moveZeroes(nums []int) {
	if len(nums) < 2 {
		return
	}

	var (
		n   = len(nums)
		cur = 0
		fz  = 0
	)

	for cur < n {
		if nums[cur] == 0 {
			// 可以跳过，因为有 fz 在前面守着
			cur++
			continue
		}

		// 说明不是 0，得看看 fz 落后没有，如果 fz 落后了，说明前面有 0，需要换位置
		if cur != fz {
			// 说明需要讲当前这个不为 0 的数往前挪到 fz 的位置
			nums[fz] = nums[cur]
			nums[cur] = 0
		}
		cur++
		fz++
	}
}

func moveZeroes2(nums []int) {
	if len(nums) < 2 {
		return
	}

	var (
		// 准备双指针
		right = len(nums) - 1
	)

	for cur := right - 1; cur >= 0; cur-- {
		if nums[cur] != 0 {
			continue
		}
		// 说明等于 0，需要移动
		for begin := cur; begin < right; begin++ {
			nums[begin] = nums[begin+1]
		}
		nums[right] = 0
		right--
	}
}
