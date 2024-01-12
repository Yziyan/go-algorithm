// @Author: Ciusyan 1/12/24

package day_30

// https://leetcode.cn/problems/remove-duplicates-from-sorted-array/description/

func removeDuplicates(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	var (
		l = len(nums)
		// 用于遍历的索引
		cur = 1
		// 已填充元素的最后一个位置
		fill = 0
	)

	for cur < l {
		if nums[cur] != nums[fill] {
			// 如果填充的最后一个位置和当前遍历的不相等，那么将其继续填充
			fill++
			nums[fill] = nums[cur]
		}

		cur++
	}

	// 遍历完后，填充到的最后一个位置就是需要的长度
	return fill + 1
}
