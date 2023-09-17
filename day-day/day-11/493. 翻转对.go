// @Author: Ciusyan 9/17/23

package day_11

// https://leetcode.cn/problems/reverse-pairs/

func reversePairs(nums []int) int {
	if nums == nil {
		return 0
	}

	return mergeSort(nums, 0, len(nums))
}

// 对 [begin, end) 进行归并排序，并且返回翻转对的数量
func mergeSort(nums []int, begin, end int) int {
	if end-begin < 2 {
		return 0
	}

	mid := begin + (end-begin)>>1

	// 最终结果 = 左边翻转对 + 右边翻转对 + 左右合并时产生的翻转对
	return mergeSort(nums, begin, mid) + mergeSort(nums, mid, end) + merge(nums, begin, mid, end)
}

// 合并 [begin, mid) 和 [mid, end)，并返回翻转对的数量
func merge(nums []int, begin, mid, end int) int {
	// 1、先处理翻转对的数量，
	count := 0
	// 代表此次寻找翻转对从哪里出发，不可以往回走，因为没必要。
	winR := mid
	// 外层对于 [begin, mid) 而言
	for i := begin; i < mid; i++ {
		// 内层对于 [mid, end) 而言
		for winR < end && nums[i] > nums[winR]<<1 {
			// 能来到这里，说明是一个翻转对，能往后走一步
			winR++
		}
		// 到了哪里 - 从哪出发 = 走了多远
		//	代表最多能往后走多少次，一路走过遇到的数字都是翻转对
		count += winR - mid
	}

	// 2、再进行归并排序
	help := make([]int, end-begin)
	// 准备几个指针
	var (
		i = 0
		l = begin
		r = mid
	)

	// 先比较左右两个序列
	for l < mid && r < end {
		// 先默认左边小
		min := nums[l]
		l++
		if min > nums[r] {
			// 说明出现了误判
			l--
			min = nums[r]
			r++
		}

		help[i] = min
		i++
	}
	// 来到这里，至少有一个序列被合并完成了，将未合并的序列添加到尾部
	for l < mid {
		// 左边未合并完成
		help[i] = nums[l]
		l++
		i++
	}
	for r < end {
		// 右边未合并完成
		help[i] = nums[r]
		r++
		i++
	}

	// 拷贝回原数组
	for _, v := range help {
		nums[begin] = v
		begin++
	}

	return count
}
