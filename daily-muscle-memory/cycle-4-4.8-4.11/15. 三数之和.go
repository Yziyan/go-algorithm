// @Author: Ciusyan 4/11/24

package cycle_4_4_8_4_11

import "sort"

// https://leetcode.cn/problems/3sum/

/**
思路重复：

我们想要找到所有三个数之和为 0 的组合。
我们该如何实现呢？
首先其实可以对这个数组进行一个排序，排序后的结果是：【负数 ... 0 ... 正数】
那么我们从前往后去计算的时候，就可以有规律可循了，
三元组是哪三个呢？[cur, left, right]
我们每次都看，这三个数是否和为 0，如果是的话，说明是一个结果。如果不是，看 left + right 的和，和 -cur 的值比起来是大了还是小了：
如果小了，就说明 left 得往后移动，因为 left 的前面一定比 left 小。
如果大了，就说明 right 得往前移动，因为 right 的后面一定比 right 大。
直到 left 和 right 相遇后，这一个范围就找完了，去下一范原来吗？围找，直至找完所有三元组。

*/

func threeSum(nums []int) [][]int {
	n := len(nums)
	res := make([][]int, 0, n)

	// 对数组排序，变成 [负数 ... 0 ... 正数]
	sort.Ints(nums)

	// 必须要有三个元素，才进去
	for cur := 0; cur < n-2; cur++ {
		if cur > 0 && nums[cur] == nums[cur-1] {
			// 说明当前数和前一个数一样，因为之前已经算过 [cur-1 ... ] 这个更大范围的结果了，就没必要算了
			continue
		}
		// 准备双指针，我们的目标是 nums[cur] + nums[left] + nums[right] = 0
		left, right := cur+1, n-1
		remain := -nums[cur] // 剩下就看 nums[left] + nums[right] 了
		for left < right {
			sumLR := nums[left] + nums[right]
			if remain == sumLR {
				// 说明得到一个答案了
				res = append(res, []int{nums[cur], nums[left], nums[right]})
				// 但是我们已经得到更大范围的了，我们将 left 和 right 周围范围更小的跳过，再去计算
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				// 但是还需要整体往中间移动
				left++
				right--
			} else if sumLR < remain {
				// 说明左右的和小了，将 left 调大一点
				left++
			} else {
				// 说明左右的和大了，将 right 调小一点
				right--
			}
		}
	}

	return res
}
