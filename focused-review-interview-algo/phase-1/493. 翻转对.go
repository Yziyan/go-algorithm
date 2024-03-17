// @Author: Ciusyan 3/17/24

package phase_1

func reversePairs(nums []int) int {

	// 合并 [begin, mid) 和 [mid, end) 两个序列，并求出过程中符合条件的数量
	merge := func(nums []int, begin, mid, end int) int {
		res := 0
		r := mid // r 不回退
		// 先看看，[begin, mid) 中，相对于 [mid, end) 能产生多少个符合条件的数
		for l := begin; l < mid; l++ {
			// 挨个检查
			for r < end && nums[l] > nums[r]<<1 {
				// 说明满足翻转对的要求
				r++
			}
			// 到达这里，说明不满足翻转对了，看看此时有几对满足翻转对的
			res += r - mid
		}

		// 下面就是归并的过程
		helpArr := make([]int, end-begin)
		r = mid
		l := begin
		hi := 0

		for l < mid && r < end {
			curMin := nums[l]
			l++
			if curMin > nums[r] {
				// 说明右边更小
				curMin = nums[r]
				r++
				l--
			}

			helpArr[hi] = curMin
			hi++
		}

		for l < mid {
			helpArr[hi] = nums[l]
			hi++
			l++
		}

		for r < end {
			helpArr[hi] = nums[r]
			hi++
			r++
		}

		for _, v := range helpArr {
			nums[begin] = v
			begin++
		}

		return res
	}

	// 求解出 nums 中，[begin, end) 中翻转对的数量
	var process func(nums []int, begin, end int) int
	process = func(nums []int, begin, end int) int {
		if end-begin < 2 {
			return 0
		}

		mid := begin + (end-begin)>>1

		return process(nums, begin, mid) + process(nums, mid, end) + merge(nums, begin, mid, end)
	}

	return process(nums, 0, len(nums))
}
