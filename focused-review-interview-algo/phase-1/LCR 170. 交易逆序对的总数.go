// @Author: Ciusyan 3/17/24

package phase_1

func reversePairs22(record []int) int {

	// 求解 record 中 [begin, end) 中逆序对的数量
	var process func(record []int, begin, end int) int
	process = func(record []int, begin, end int) int {
		if end-begin < 2 {
			// 都小于两个元素了
			return 0
		}
		mid := begin + (end-begin)>>1

		// 其中逆序对的数量就是：[begin, mid) 的 + [mid, end) 的 + 合并左右两个序列产生的
		return process(record, begin, mid) + process(record, mid, end) + reversePairsMerge(record, begin, mid, end)
	}

	// 求解 [0, len) 中逆序对的数量
	return process(record, 0, len(record))
}

// 对 nums 中 [begin, mid) 和 [mid, end) 两个序列进行降序合并，并且返回其中逆序对的数量
func reversePairsMerge(nums []int, begin, mid, end int) int {
	res := 0
	helpArr := make([]int, end-begin)

	var (
		l  = begin
		r  = mid
		hi = 0
	)

	for l < mid && r < end {
		// 默认右边的数最大，就不会产生逆序对了
		curMax := nums[r]
		r++
		if curMax < nums[l] {
			// 说明左边的要大，会产生 end-r 个逆序对
			curMax = nums[l]
			l++
			r--

			// 相加上产生的逆序对
			res += end - r
		}

		helpArr[hi] = curMax
		hi++
	}

	for l < mid {
		// 说明左边还没有合并完
		helpArr[hi] = nums[l]
		hi++
		l++
	}

	for r < end {
		// 说明右边还没合并完
		helpArr[hi] = nums[r]
		hi++
		r++
	}

	// 作用于源数组
	for _, v := range helpArr {
		nums[begin] = v
		begin++
	}

	return res
}
