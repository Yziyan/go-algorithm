// @Author: Ciusyan 3/17/24

package phase_1

import "math/rand"

func QuickSortArray(nums []int) []int {

	// 返回轴点元素 nums[begin]，的左右边界
	getPivot := func(nums []int, begin, end int) (int, int) {
		if end-begin == 1 {
			// 说明只有一个元素
			return begin, begin
		}

		var (
			pivot = nums[begin]
			l     = begin + 1
			cur   = l
			r     = end - 1
		)

		for cur <= r {
			if nums[cur] < pivot {
				// 说明比 pivot 小，放左边
				nums[cur], nums[l] = nums[l], nums[cur]
				l++
				cur++
			} else if nums[cur] == pivot {
				// 相等，先待定
				cur++
			} else {
				// 说明比轴点大，确定在轴点的后面了
				nums[cur], nums[r] = nums[r], nums[cur]
				r-- // 末尾确定了，但是被换过来的元素也要检查
			}
		}

		// 来到这里，已经全部划分完毕了，将轴点交换到合适位置
		nums[begin], nums[l-1] = nums[l-1], nums[begin]

		// 此时 l-1 才是轴点的左边界了，右边界刚好是 r
		return l - 1, r
	}

	// 对 num 的 [begin, end) 进行快速排序
	var sort func(nums []int, begin, end int)
	sort = func(nums []int, begin, end int) {
		if end-begin < 2 {
			// 说明没有两个元素
			return
		}

		// 否则随机选取一个范围在 [begin, end) 之间的轴点
		randIdx := rand.Intn(end-begin) + begin
		nums[begin], nums[randIdx] = nums[randIdx], nums[begin]
		// 获取轴点的范围
		l, r := getPivot(nums, begin, end)

		// 对轴点左右进行快排
		sort(nums, begin, l)
		sort(nums, r+1, end)
	}

	sort(nums, 0, len(nums))
	return nums
}
