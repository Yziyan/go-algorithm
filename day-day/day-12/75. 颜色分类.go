// @Author: Ciusyan 9/18/23

package day_12

// https://leetcode.cn/problems/sort-colors/description/

func sortColors(nums []int) {
	if nums == nil {
		return
	}

	// 准备几个指针
	var (
		cur = 0
		l   = 0
		r   = len(nums) - 1
	)

	// 1 是参考值，所有数字都需要与参考值基准一遍
	for cur <= r {
		if nums[cur] < 1 {
			// 小于 1，cur 和 l 交换，并且都往后走
			nums[cur], nums[l] = nums[l], nums[cur]
			cur++
			l++
		} else if nums[cur] == 1 {
			cur++
		} else {
			// 大于 1，cur 和 r 交换，只将 r 往前走，cur 继续基准，因为刚从后面过来之前没有比较过
			nums[cur], nums[r] = nums[r], nums[cur]
			r--
		}
	}
}
