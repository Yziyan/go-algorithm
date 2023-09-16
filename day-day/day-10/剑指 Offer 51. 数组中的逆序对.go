// @Author: Ciusyan 9/16/23

package day_10

// https://leetcode.cn/problems/shu-zu-zhong-de-ni-xu-dui-lcof/

func reversePairs(nums []int) int {
	if nums == nil {
		return 0
	}

	return mergeSort(nums, 0, len(nums))
}

// nums [begin, end) 中逆序对的数量，但是我排序按照逆序排
func mergeSort(nums []int, begin, end int) int {
	if end-begin < 2 {
		return 0
	}

	mid := begin + (end-begin)>>1

	// 返回左边逆序对的数量 + 右边逆序对的数量 + 整体合并逆序对的数量
	return mergeSort(nums, begin, mid) + mergeSort(nums, mid, end) + mergeAndCount(nums, begin, mid, end)
}

// 合并 [begin, mid) 和 [mid, end) 两个数组，按照降序排列，并且合并时计算有多少个逆序对
func mergeAndCount(nums []int, begin, mid, end int) int {

	help := make([]int, end-begin)
	count := 0
	// 准备几个指针
	var (
		i = 0
		l = begin
		r = mid
	)

	// 先比较合并
	for l < mid && r < end {

		// 先默认右边大
		max := nums[r]
		r++

		if max < nums[l] {
			// 说明误判了
			r--

			// 说明右边有 (end - r) 个数，比 nums[l] 小，累加到逆序对数量中
			count += end - r

			max = nums[l]
			l++
		}

		// 更新 help
		help[i] = max
		i++
	}

	// 来到这里，肯定有一边合并完了，将剩下的一边放置到数组中
	for l < mid {
		// 说明左边没合并完
		help[i] = nums[l]
		l++
		i++
	}
	for r < end {
		// 说明右边没合并完
		help[i] = nums[r]
		r++
		i++
	}

	// 将 help 设置到 nums[begin, end) 上
	for _, v := range help {
		nums[begin] = v
		begin++
	}

	return count
}
