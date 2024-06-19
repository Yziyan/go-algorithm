// @Author: Ciusyan 6/19/24

package phase3

// https://leetcode.cn/problems/trapping-rain-water/
func trap(height []int) int {
	var (
		l = 0
		r = len(height) - 1

		water    = 0
		maxLower = 0
	)

	for l <= r {
		lower := height[l]
		l++
		if lower > height[r] {
			lower = height[r]
			r--
			l--
		}

		if lower > maxLower {
			maxLower = lower
		}

		water += maxLower - lower
	}

	return water
}
