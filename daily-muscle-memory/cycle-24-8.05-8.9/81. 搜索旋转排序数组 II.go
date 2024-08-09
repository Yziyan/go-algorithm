// @Author: Ciusyan 2024/8/9

package cycle_24_8_05_8_9

// https://leetcode.cn/problems/search-in-rotated-sorted-array-ii/

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
