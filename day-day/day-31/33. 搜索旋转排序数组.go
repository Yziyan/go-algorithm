// @Author: Ciusyan 1/17/24

package day_31

// https://leetcode.cn/problems/search-in-rotated-sorted-array/description/

func search1(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}

	var (
		// 在 [begin, end] 上，能二分就尽量的二分
		begin = 0
		end   = len(nums) - 1
		mid   = 0
	)

	// 只有一个数的时候，也得看看是否是 target
	for begin <= end {
		// 先计算出中点位置
		mid = begin + ((end - begin) >> 1)
		if nums[mid] == target {
			// 如果中点直接找到了，运气也太好了
			return mid
		}

		// 找不到，分区间来二分，因为原来的有序数组被旋转过了，
		// 可能需要分区间了，分为什么区间呢？[begin, mid] 和 [mid+1, end]
		if nums[begin] == nums[mid] && nums[mid] == nums[end] {
			// 说明 begin mid end 三个位置的数相等
			for begin != mid && nums[begin] == nums[mid] {
				// 那么让 begin 先走到第一个不和 mid 相等的数，
				// 因为上面已经确认过了， mid 不可能是结果
				begin++
			}

			// 来到这里，有两种情况：
			// 	1.begin == mid
			// 	2.nums[begin] != nums[mid]
			if begin == mid {
				// 如果是第一种情况，那么如果有答案，也只可能是在 [mid+1, end] 这个区间了
				begin = mid + 1
				// 那么下一次去 [mid+1, end] 上查找 target
				continue
			}
		}

		// 来到这里，说明 begin、mid、end 三个位置的数，不都相等。
		if nums[begin] != nums[mid] {
			// 说明 begin != mid ? end
			if nums[mid] > nums[begin] {
				// 说明 [begin, mid] 是有序的，断点肯定在 [mid+1, end]，比如 [3, 3, 4, (5), |1, 2, 2]
				if target >= nums[begin] && target < nums[mid] {
					// 说明 target 只可能在 [begin, mid-1] 中，在这里面二分
					end = mid - 1
				} else {
					// 说明 target 只可能在 [mid+1, end] 中
					begin = mid + 1
				}
			} else {
				// 和上面对称，说明 [begin, mid] 是无序的，断点肯定在这里面，比如 [ |3, 4, 5, (1), 2, 2, 3]
				if target > nums[mid] && target <= nums[end] {
					// 说明 target 只可能在 [mid+1, end] 中
					begin = mid + 1
				} else {
					// 说明 target 只可能在 [begin, mid-1] 中
					end = mid - 1
				}
			}

		} else {
			// 说明 begin = mid != end，那么有两种情况：
			// 1.[mid, end] 上是有序的，断点只可能在前面 比如 [ |2, 2, (2), 3, 3, 4]
			// 2.[mid, end] 上是无序的，断点在这里面，比如 [3, 3, 3, (3), 4, 5, 1]
			// 但无论是哪种情况，答案都只可能在后面，因为 [begin ... mid] 都是相等的，并且不等于 target
			begin = mid + 1
		}

	}

	return -1
}

// https://leetcode.cn/problems/search-in-rotated-sorted-array-ii/ ，如果是这个题，记得将返回值改一下
func search(nums []int, target int) bool {
	if nums == nil || len(nums) == 0 {
		return false
	}

	var (
		// 在 [begin, end] 上，能二分就尽量的二分
		begin = 0
		end   = len(nums) - 1
		mid   = 0
	)

	// 只有一个数的时候，也得看看是否是 target
	for begin <= end {
		// 先计算出中点位置
		mid = begin + ((end - begin) >> 1)
		if nums[mid] == target {
			// 如果中点直接找到了，运气也太好了
			return true
		}

		// 找不到，分区间来二分，因为原来的有序数组被旋转过了，
		// 可能需要分区间了，分为什么区间呢？[begin, mid] 和 [mid+1, end]
		if nums[begin] == nums[mid] && nums[mid] == nums[end] {
			// 说明 begin mid end 三个位置的数相等
			for begin != mid && nums[begin] == nums[mid] {
				// 那么让 begin 先走到第一个不和 mid 相等的数，
				// 因为上面已经确认过了， mid 不可能是结果
				begin++
			}

			// 来到这里，有两种情况：
			// 	1.begin == mid
			// 	2.nums[begin] != nums[mid]
			if begin == mid {
				// 如果是第一种情况，那么如果有答案，也只可能是在 [mid+1, end] 这个区间了
				begin = mid + 1
				// 那么下一次去 [mid+1, end] 上查找 target
				continue
			}
		}

		// 来到这里，说明 begin、mid、end 三个位置的数，不都相等。
		if nums[begin] != nums[mid] {
			// 说明 begin != mid ? end
			if nums[mid] > nums[begin] {
				// 说明 [begin, mid] 是有序的，断点肯定在 [mid+1, end]，比如 [3, 3, 4, (5), |1, 2, 2]
				if target >= nums[begin] && target < nums[mid] {
					// 说明 target 只可能在 [begin, mid-1] 中，在这里面二分
					end = mid - 1
				} else {
					// 说明 target 只可能在 [mid+1, end] 中
					begin = mid + 1
				}
			} else {
				// 和上面对称，说明 [begin, mid] 是无序的，断点肯定在这里面，比如 [ |3, 4, 5, (1), 2, 2, 3]
				if target > nums[mid] && target <= nums[end] {
					// 说明 target 只可能在 [mid+1, end] 中
					begin = mid + 1
				} else {
					// 说明 target 只可能在 [begin, mid-1] 中
					end = mid - 1
				}
			}

		} else {
			// 说明 begin = mid != end，那么有两种情况：
			// 1.[mid, end] 上是有序的，断点只可能在前面 比如 [ |2, 2, (2), 3, 3, 4]
			// 2.[mid, end] 上是无序的，断点在这里面，比如 [3, 3, 3, (3), 4, 5, 1]
			// 但无论是哪种情况，答案都只可能在后面，因为 [begin ... mid] 都是相等的，并且不等于 target
			begin = mid + 1
		}

	}

	return false
}
