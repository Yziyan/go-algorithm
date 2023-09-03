// @Author: Ciusyan 2023/9/3

package day_4

// https://leetcode.cn/problems/merge-sorted-array/

func merge(nums1 []int, m int, nums2 []int, n int) {
	if nums1 == nil || nums2 == nil {
		return
	}

	// 倒着合并
	var (
		// nums1 和 nums2 目前的索引
		n1i = m - 1
		n2i = n - 1

		// 目前能存放的位置
		cur = m + n - 1
	)

	// 第二个数组合并完了，就可以停了
	for n2i >= 0 {

		if n1i >= 0 && nums1[n1i] > nums2[n2i] {
			// 说明 nums1 的必须放后面去
			nums1[cur] = nums1[n1i]
			n1i--
		} else {
			// nums2 放到合适位置，因为我们要尽可能让 n2i 消耗完
			nums1[cur] = nums2[n2i]
			n2i--
		}

		cur--
	}
}
