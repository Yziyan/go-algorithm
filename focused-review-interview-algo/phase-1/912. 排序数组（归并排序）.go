// @Author: Ciusyan 3/17/24

package phase_1

func sortArray(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	// 对 nums 的 [begin, end) 进行排序
	var sort func(nums []int, begin, end int)
	sort = func(nums []int, begin, end int) {
		if end-begin < 2 {
			// 说明元素都不足两个，排什么排
			return
		}

		// 计算 mid
		mid := begin + (end-begin)>>1
		sort(nums, begin, mid) // 对 [begin, mid) 进行排序
		sort(nums, mid, end)   // 对 [mid, end) 进行排序
		// 对排序后的两个序列，进行合并
		sortArrayMerge(nums, begin, mid, end)
	}

	// 对序列的 [0, len) 进行排序
	sort(nums, 0, len(nums))
	return nums
}

// 对 nums 的 [begin, mid) 和 [mid, end) 进行合并
func sortArrayMerge(nums []int, begin, mid, end int) {
	// 中转数组
	helpArr := make([]int, end-begin)

	var (
		l  = begin
		r  = mid
		hi = 0
	)

	for l < mid && r < end {
		// 左右两边都不越界时，才来比较
		cur := nums[l]
		l++
		if cur > nums[r] {
			// 说明右边还要小
			cur = nums[r]
			r++
			l--
		}

		helpArr[hi] = cur
		hi++
	}

	// 到达这里，说明至少有一边合并完成了，将剩下的一边放入对应数组
	if l < mid {
		// 说明左边还没合并完成
		for l < mid {
			helpArr[hi] = nums[l]
			hi++
			l++
		}
	} else {
		// 说明右边还没合并完成
		for r < end {
			helpArr[hi] = nums[r]
			hi++
			r++
		}
	}

	// 然后更新会原始序列
	for i, v := range helpArr {
		nums[begin+i] = v
	}
}
