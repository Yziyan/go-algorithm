// @Author: Ciusyan 9/21/23

package day_14

// https://leetcode.cn/problems/trapping-rain-water/

func trap(height []int) int {
	if height == nil || len(height) < 3 {
		return 0
	}

	water := 0
	lowerMax := 0

	var (
		left  = 0
		right = len(height) - 1
	)

	for left < right {
		// 先假设左边更矮
		lower := height[left]
		left++
		if lower > height[right] {
			// 说明右边更低，刚刚误判了
			left--
			// 以右边对齐
			lower = height[right]
			right--
		}

		// 来到这里，尝试更新一直以来，左右两边最大值
		if lower > lowerMax {
			lowerMax = lower
		}
		// 积累雨水
		water += lowerMax - lower
	}

	return water
}
