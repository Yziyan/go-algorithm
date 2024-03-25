// @Author: Ciusyan 3/26/24

package cycle_3_3_24_3_29

// https://leetcode.cn/problems/maximum-product-subarray/description/

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var (
		// 前一个数为结尾位置时，所得到子数组乘积的最大值和最小值
		preMax = nums[0]
		preMin = nums[0]

		maxRes = nums[0]
	)

	for i := 1; i < len(nums); i++ {
		// 有三种情况，能得到最大值
		p1 := nums[i]          // 自己
		p2 := nums[i] * preMax // 自己 * 前面的最大值
		p3 := nums[i] * preMin // 自己 * 前面的最小值

		curMax := max(p1, p2, p3)
		curMin := min(p1, p2, p3)

		// 看看当前能否将结果推大
		maxRes = max(maxRes, curMax)
		preMax = curMax
		preMin = curMin
	}

	return maxRes
}
