// @Author: Ciusyan 1/14/24

package day_30

// https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/

func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	if nums == nil || len(nums) == 0 {
		return res
	}
	n := len(nums)
	// 利用二分查找，找出第一个 =target 的数
	l := lowerRightIdx(nums, target)
	if l == n || nums[l] != target {
		// 说明数组中没有 target 这个数
		return res
	}

	// 到达这里，说明 L 位置已经确认了，
	res[0] = l

	if l == n-1 {
		// 说明第一个等于 target 的就是数组的最后一个元素，那么最右也是
		res[1] = l
	} else {
		// 否则得再调用上述方法，找到第一个 = target+1 的数，减去1 就是对应 r 的索引
		res[1] = lowerRightIdx(nums, target+1) - 1
	}

	return res
}

// 返回 nums 中，第一个 = target 的索引，使用二分查找
func lowerRightIdx(nums []int, target int) int {
	var (
		l = 0
		r = len(nums)
	)
	// 查找的范围是：[l, r)
	for l < r {
		// 取中间位置
		mid := l + ((r - l) >> 1)
		if nums[mid] < target {
			// 说明比目标小，继续往右查找
			l = mid + 1
		} else {
			// 说明当前中间位置元素 >= target，
			// 大于的时候肯定要往左走，这个好理解
			// 但当等于的时候，为啥也要往左走呢？因为有可能等于了也不是第一个数，所以继续往左走搜索
			r = mid
		}
	}

	return l
}
