// @Author: Ciusyan 5/19/24

package cycle_12_5_18_5_22

// https://leetcode.cn/problems/merge-sorted-array/
func merge(nums1 []int, m int, nums2 []int, n int) {
	l := m - 1
	r := n - 1
	i := m + n - 1

	// 以右边为基准，只要右边合并完成即说明整体合完了
	for r >= 0 {
		// 所以优先合并右边
		if l >= 0 && nums1[l] > nums2[r] {
			// 左边的大
			nums1[i] = nums1[l]
			l--
		} else {
			// 前面的要大，等也先放后面的
			nums1[i] = nums2[r]
			r--
		}
		i--
	}
}
