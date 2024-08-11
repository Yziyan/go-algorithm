// @Author: Ciusyan 2024/8/9

package cycle_24_8_05_8_9

// https://leetcode.cn/problems/search-in-rotated-sorted-array-ii/

/**
思路重复：
比起旋转数组一，这里多了可能有重复元素这个限制。

那么只要我们在处理的过程中，能过跳过这样的限制，即可按照一一样的方式处理。

如何处理呢？
也是开始二分搜索，
先看是否 mid 是值，如果不是，我们先看 l, r, mid  三个位置是否都一样，如果都一样，我们先保证，三个数不都一样。

那么当跳过后，就看断点在前面还是后面了。
如果 nums[mid] >= nums[l] 说明答案可能在前面，可以在前面就行二分，前面有序。
反之则是后面有序

*/

func search2(nums []int, target int) bool {
	l, r, mid := 0, len(nums)-1, 0

	for l <= r {
		mid = (l + r) / 2
		if nums[mid] == target {
			return true
		}

		// 说明还没找到，先确保 l r mid 三个位置不都一样
		if nums[l] == nums[mid] && nums[mid] == nums[r] {
			// 说明三个数一样
			for l != mid && nums[l] == nums[r] {
				l++
			}

			// 来到这里，只有两种情况、
			// 1. l == mid，说明答案只可能在 mid+1~r 上
			// 2. nums[l] != nums[r], 说明达到了目的，不都一样，去下面统一处理即可
			if l == mid {
				// 这种情况
				l = mid + 1
				continue
			}
		}

		if nums[mid] >= nums[l] {
			// 说明 l...mid 可以进行二分
			if target < nums[mid] && target >= nums[l] {
				// 说明在前半段
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			// 说明 mid...r 可以进行二分
			if target <= nums[r] && target > nums[mid] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}

	return false
}
