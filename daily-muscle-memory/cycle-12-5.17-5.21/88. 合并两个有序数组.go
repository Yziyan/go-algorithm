// @Author: Ciusyan 5/19/24

package cycle_12_5_17_5_21

// https://leetcode.cn/problems/merge-sorted-array/

/*
*
思路重复：
这就是归并排序中，归并的流程。
如何将两个有序数组合并为一个有序数组。
使用多指针，从两个数组的后方开始合并。
挨个比较两个指向位置的大小，将大的放置到待放置的位置。然后进行下一次比较。
大体思路就是上面那样，但是呢？我们有两个小技巧，比如说：
核心看看右边的数组有无合并结束，如果合并结束，那么整个合并就结束了。
所以当等于的时候，也优先取右边的数字。
*/
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
