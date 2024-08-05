// @Author: Ciusyan 2024/8/5

package cycle_24_8_05_8_9

// https://leetcode.cn/problems/search-in-rotated-sorted-array/

/**
这题的数，整体还是有序的，能不能用二分搜索呢？
其实是可以的，但因为中间有断点，所以不能直接使用二分查找。
只能在断点的前后两段进行二分，
那么核心就是看，tgt 属于哪一段区间。
如果属于前一段，那么去前一段二分即可。如果属于后一段，同样去后一段二分即可。
*/

func search2(nums []int, target int) int {
	l, mid, r := 0, 0, len(nums)-1

	for l <= r {
		mid = (l + r) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] >= nums[l] {
			// 说明前半段有序
			if target < nums[mid] && target >= nums[l] {
				// 说明再前半段
				r = mid - 1
			} else {
				l = mid + 1
			}

		} else {
			if target <= nums[r] && target > nums[mid] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}

	return -1
}
func search(nums []int, target int) int {
	l, r, mid := 0, len(nums)-1, 0

	for l <= r {
		mid = (l + r) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] >= nums[l] {
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
