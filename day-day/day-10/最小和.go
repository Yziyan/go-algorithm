// @Author: Ciusyan 9/16/23

package day_10

// SmallSum 什么是最小和呢？每一个数左边比他小的数的和的总和。比如 nums = [1, 3, 7, 2, 1, 8, 5]
//
//	1: 0
//	3: 1
//	7: 4
//	2: 1
//	1: 0
//	8: 14
//	5: 7
//	SmallSum = 0+1+4+1+0+14+7
func SmallSum(nums []int) int {
	if nums == nil {
		return 0
	}

	// 求解 [0, len) 的最小和
	return process(nums, 0, len(nums))
}

// 返回 nums 中 [begin, end) 的 SmallSum
func process(nums []int, begin, end int) int {
	if end-begin < 2 {
		return 0
	}

	mid := begin + (end-begin)>>1

	// 左边的最小和 + 右边的最小和 + 合并时产生的最小和
	return process(nums, begin, mid) + process(nums, mid, end) + merge(nums, begin, mid, end)
}

// 合并 nums 中 [begin, mid) 和 [mid, end)，并返回合并时产生的最小和
func merge(nums []int, begin, mid, end int) int {

	help := make([]int, end-begin)
	// 最小和
	smallSum := 0

	// 准备几个指针
	var (
		i = 0
		l = begin
		r = mid
	)

	// 只要有一边合并完成了，就退出
	for l < mid && r < end {
		// 先默认右边小
		min := nums[r]
		r++
		// 如果相等，也先挪动后面的内容
		if min > nums[l] {
			min = nums[l]

			// 先还原误判的指针
			r--

			// 需要累加最小和，因为此时找到了 (end - r) 个比 nums[l] 大的数
			smallSum += nums[l] * (end - r)

			l++
		}

		// 放入小的
		help[i] = min
		i++
	}

	// 来到这里，肯定有一边退出了，将未合并的部分直接写如 help 中
	for l < mid {
		// 前面没合并完
		help[i] = nums[l]
		l++
		i++
	}

	for r < end {
		// 后面没合并完
		help[i] = nums[r]
		r++
		i++
	}

	// 最终将合并的结果，放入 [begin, end)
	for _, v := range help {
		nums[begin] = v
		begin++
	}

	return smallSum
}
