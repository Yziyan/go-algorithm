// @Author: Ciusyan 3/27/24

package cycle_3_3_24_3_29

// https://leetcode.cn/problems/maximum-subarray/

/**
思路重复：
比昨天那个题还简单，思路类似的。
都是：将 nums[0...i] 去求解出一个最大的子数组和，要求必须以 nums[i] 结尾。
	然后利用前一个结果，去推 nums[0...i+1] 的结果。
有两种情况：
	1.就是 nums[i+1]
	2.还要加上前面得到的最大值，但是要求前面的结果，对现在有增益。
*/

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var (
		preMax = nums[0]

		maxRes = nums[0]
	)

	for i := 1; i < len(nums); i++ {
		p1 := nums[i]
		if preMax > 0 {
			p1 += preMax
		}

		maxRes = max(maxRes, p1)
		preMax = p1
	}

	return maxRes
}

func maxSubArray2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var (
		preMax = nums[0]

		maxRes = nums[0]
	)

	for i := 1; i < len(nums); i++ {
		// 第一种情况，就是自己
		p1 := nums[i]
		// 第二种情况，包括前面的 preMax
		if preMax > 0 {
			// 说明前面的有收益
			p1 += preMax
		}

		maxRes = max(maxRes, p1)
		preMax = p1
	}

	return maxRes
}
