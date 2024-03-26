// @Author: Ciusyan 3/26/24

package cycle_3_3_24_3_29

// https://leetcode.cn/problems/maximum-product-subarray/description/

/**
思路重复：
要求子数组最大的乘积，核心下面几点：
1.从头开始，把每一位当做结尾，求解出以当时的最后一位结尾，得到的最大子数组乘积。
2.扩充下一位的时候，有三种可能：
	2.1自己最大（正负得负）
	2.2自己和之前的最大值相乘最大（正正得正）
	2.3自己和之前的最小值相乘最大（负负得正）
3.所以每次遍历，需要保留前一个数为结尾时，得到的最大、最小值
*/

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var (
		preMax = nums[0]
		preMin = nums[0]

		maxRes = nums[0]
	)

	for i := 1; i < len(nums); i++ {
		p1 := nums[i]
		p2 := nums[i] * preMax
		p3 := nums[i] * preMin

		curMax := max(p1, p2, p3)
		curMin := min(p1, p2, p3)

		maxRes = max(maxRes, curMax)

		preMax = curMax
		preMin = curMin
	}

	return maxRes
}

func maxProduct2(nums []int) int {
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
