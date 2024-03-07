// @Author: Ciusyan 3/7/24

package day_34

// https://leetcode.cn/problems/maximum-product-subarray/

func maxProduct(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	var (
		preMax = nums[0] // 前一个位置的最大值
		preMin = nums[0] // 前一个位置的最小值

		maxRes = nums[0] // 结果
	)

	// 挨个位置求解，通过 cur-1 位置的两个 pre 信息，求解出 cur 位置的答案
	for cur := 1; cur < len(nums); cur++ {
		// 有三种情况
		p1 := nums[cur]          // 就是他自己 eg: [0,| 3]
		p2 := preMax * nums[cur] // 是前一位置的最大值 * 自己 eg: [2, 3,| 5]
		p3 := preMin * nums[cur] // 是前一位置的最小值 * 自己 eg: [3, -8,| -9]

		curMax := max(p1, p2, p3)
		curMin := min(p1, p2, p3)

		maxRes = max(maxRes, curMax)

		// 记得去下一个位置之前，要把当前的值给前一个位置
		preMax = curMax
		preMin = curMin
	}

	return maxRes
}
