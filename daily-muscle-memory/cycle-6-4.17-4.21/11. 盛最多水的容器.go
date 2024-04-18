// @Author: Ciusyan 4/19/24

package cycle_6_4_17_4_21

// https://leetcode.cn/problems/container-with-most-water/

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1

	res := 0
	for left < right {
		lower := height[left]
		if lower > height[right] {
			lower = height[right]
		}

		curRes := (right - left) * lower
		if curRes > res {
			res = curRes
		}

		// 看看左右两边，如果都比当前柱子最低的柱子还低，就没必要算了
		for left < right && height[left] <= lower {
			left++
		}
		for left < right && height[right] <= lower {
			right--
		}
	}

	return res
}
