// @Author: Ciusyan 1/22/24

package day_32

// https://leetcode.cn/problems/merge-sorted-array/

func merge(nums1 []int, m int, nums2 []int, n int) {

	var (
		idx  = len(nums1) - 1 // 需要插入位置的索引
		cur1 = m - 1          // 指向 nums1 数组当前的元素
		cur2 = n - 1          // 指向 nums2 数组当前的元素
	)

	// 当 nums2 合并完了，就说明合并完了
	for cur2 >= 0 {
		if cur1 >= 0 && nums1[cur1] > nums2[cur2] {
			// 说明前面的数组大，放到 idx 位置
			nums1[idx] = nums1[cur1]
			cur1--
		} else {
			// 说明后面的数组大，
			// 相等也放在这里，优先合并完后面的数组，就可以快点结束循环了
			nums1[idx] = nums2[cur2]
			cur2--
		}
		// 填完当前位置后，索引前移
		idx--
	}
}
