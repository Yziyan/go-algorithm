// @Author: Ciusyan 4/25/24

package cycle_7_4_22_4_26

// https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/description/

/**
思路重复：
既然是有序数组，我们看到这个字眼，应该就得想到：二分搜索。
那么这个题，应该如何二分呢？
1.二分去查找这个数组中，target 出现的第一个位置。
在二分的时候，遇到小于 target 的，往后走；否则往前走，一直走到最后。
2.然后看看找到的位置是否是合法的 target。是的话就去查找 target+1 在数组中的位置。
然后 -1，就是最后一个位置。


*/

func searchRange(nums []int, target int) []int {

	findFirstIdx := func(nums []int, tgt int) int {
		l, r := 0, len(nums)
		for l < r {
			mid := l + ((r - l) >> 1)
			if nums[mid] < tgt {
				l = mid + 1
			} else {
				r = mid
			}
		}

		return l
	}

	n := len(nums)
	l := findFirstIdx(nums, target)

	res := []int{-1, -1}
	if l == n || nums[l] != target {
		return res
	}

	res[0] = l
	if l+1 == n || nums[l+1] != target {
		res[1] = l
	} else {
		res[1] = findFirstIdx(nums, target+1) - 1
	}

	return res
}

func searchRange1(nums []int, target int) []int {
	res := []int{-1, -1}
	if len(nums) == 0 {
		return res
	}

	n := len(nums)

	// 先找到第一个位置
	l := findFirstTgt(nums, target)
	if l == n || nums[l] != target {
		// 说明不存在 target
		return res
	}

	// 来到这里，说明存在 l 了
	res[0] = l
	if l == n-1 || nums[l+1] != target {
		// 说明只有一个 target
		res[1] = l
	} else {
		// 说明有多个，去二分查找 target+1，它的位置减一，就是最后一个 target 的位置
		res[1] = findFirstTgt(nums, target+1) - 1
	}

	return res
}

// 返回 target 在 nums 中，第一个位置的索引，
func findFirstTgt(nums []int, target int) int {
	l, r := 0, len(nums)

	// 在 [l ... r) 上，进行二分查找
	for l < r {
		// 取出中点
		mid := l + ((r - l) >> 1)
		if nums[mid] < target {
			// 说明小了，往后靠
			l = mid + 1
		} else {
			// 大于等于 target，中间往前靠。即使是等于，也不能返回，因为我们要找第一个位置
			r = mid
		}
	}

	return l
}
