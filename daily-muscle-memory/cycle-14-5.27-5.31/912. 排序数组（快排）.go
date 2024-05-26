// @Author: Ciusyan 5/27/24

package cycle_14_5_27_5_31

import "math/rand"

// https://leetcode.cn/problems/sort-an-array/

func sortArray(nums []int) []int {
	pivot := func(nums []int, l, r int) (lp, rp int) {
		if r-l == 1 {
			// 说明只有一个元素
			return l, r
		}
		// 随机一个元素到 nums[l] 作为轴点
		randIdx := rand.Intn(r-l) + l
		nums[l], nums[randIdx] = nums[randIdx], nums[l]

		pivotIdx := l
		target := nums[pivotIdx]
		l++
		r--

		// 现在 nums[l] 就是 target，将 nums 分为三段：小于 target | 等于 target | 大于 target
		for cur := l; cur <= r; {
			if nums[cur] < target {
				// 说明 cur 应该在 target 左边
				nums[cur], nums[l] = nums[l], nums[cur]
				cur++
				l++
			} else if nums[cur] == target {
				// 说明相等，cur 往后走即可
				cur++
			} else {
				// 说明 cur 应该在 target 右边
				nums[cur], nums[r] = nums[r], nums[cur]
				r--
				// cur 不能动，因为换过来的元素，也还得进行比较
			}
		}

		// 将轴点放在合适的位置
		nums[pivotIdx], nums[l-1] = nums[l-1], nums[pivotIdx]

		return l - 1, r
	}

	// 对 [l, r) 进行快速排序
	var quickSort func(nums []int, l, r int)
	quickSort = func(nums []int, l, r int) {
		if r-l < 2 {
			return
		}

		// 寻找轴点的左右两端
		lp, rp := pivot(nums, l, r)
		quickSort(nums, l, lp)
		quickSort(nums, rp+1, r)
	}

	quickSort(nums, 0, len(nums))
	return nums
}
