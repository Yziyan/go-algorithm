// @Author: Ciusyan 5/27/24

package cycle_14_5_27_5_31

// https://leetcode.cn/problems/sort-an-array/

func sortArray(nums []int) []int {
	merge := func(nums []int, l, mid, r int) {
		// 申请辅助数组，用来装合并后的结果
		helpArr := make([]int, r-l)

		li, ri := l, mid
		hi := 0

		// 左右都有元素没合并时，就直接合并
		for li != mid && ri != r {
			if nums[li] < nums[ri] {
				helpArr[hi] = nums[li]
				hi++
				li++
			} else {
				helpArr[hi] = nums[ri]
				hi++
				ri++
			}
		}

		// 至少有一边是合并完成了
		for li != mid {
			helpArr[hi] = nums[li]
			li++
			hi++
		}

		for ri != r {
			helpArr[hi] = nums[ri]
			ri++
			hi++
		}

		// 将排序结果变更到原序列中
		hi = 0
		for l < r {
			nums[l] = helpArr[hi]
			l++
			hi++
		}
	}

	// 对 nums 的 [l, r) 范围，进行归并排序
	var sort func(nums []int, l, r int)
	sort = func(nums []int, l, r int) {
		if r-l < 2 {
			return
		}

		// 先计算出 l ... r 的中点
		mid := l + ((r - l) >> 1)
		sort(nums, l, mid)     // 左边进行归并
		sort(nums, mid, r)     // 右边进行归并
		merge(nums, l, mid, r) // 左右子序列进行合并
	}

	sort(nums, 0, len(nums))
	return nums
}
