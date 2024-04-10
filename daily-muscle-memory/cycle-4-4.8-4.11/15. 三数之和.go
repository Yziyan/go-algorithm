// @Author: Ciusyan 4/11/24

package cycle_4_4_8_4_11

import "sort"

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
