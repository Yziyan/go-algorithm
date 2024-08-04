// @Author: Ciusyan 2024/8/5

package cycle_24_8_05_8_9

// https://leetcode.cn/problems/search-in-rotated-sorted-array/

func search(nums []int, target int) int {
	l, r, mid := 0, len(nums)-1, 0

	for l <= r {
		mid = (l + r) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] > nums[l] {
			// 说明断点在后面，前面可以二分
			if target < nums[mid] && target >= nums[l] {
				// 说明 target ∈ [l, mid)
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			// 说明断点在前面，后面可以二分
			if target <= nums[r] && target > nums[mid] {
				// 说明 target ∈ (mid, r]
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}

	return -1
}
