// @Author: Ciusyan 2023/9/3

package day_4

// https://leetcode.cn/problems/trapping-rain-water/

// 解法 3，双指针优雅的写法
func trap(height []int) int {
	if height == nil || len(height) < 3 {
		return 0
	}

	var (
		left  = 0
		right = len(height) - 1

		// 较低的最大值
		lowerMax = 0
	)

	water := 0
	for left < right {
		lower := 0
		if height[left] > height[right] {
			// 说明右边低，不用考虑左边了
			lower = height[right]
			// 右边往中间靠
			right--
		} else {
			// 说明左边低，不用考虑右边了
			lower = height[left]
			// 左边往中间靠
			left++
		}
		// 先尝试更新较低值的最大值
		if lower > lowerMax {
			lowerMax = lower
		}

		// 那么能接的水：
		water += lowerMax - lower
	}

	return water
}

// 解法 2，双指针解法
func trap2(height []int) int {
	if height == nil || len(height) < 3 {
		return 0
	}

	var (
		// 柱子的索引
		left  = 1
		right = len(height) - 2
		// 左右两边的最大值
		leftMax  = 0
		rightMax = 0
	)

	water := 0
	for left <= right {
		if height[left-1] > height[right+1] {
			// 右边低一些，说明不用考虑左边了
			//  从右边看，先尝试更新右边柱子的最大值
			if height[right+1] > rightMax {
				rightMax = height[right+1]
			}
			// 当前柱子比右边最高的柱子还低才能接水
			diff := rightMax - height[right]
			if diff > 0 {
				water += diff
			}

			// 接完后往中间靠
			right--
		} else {
			// 说明左边低，不用考虑右边了
			//  从左边看，先更新左边柱子的最大值
			if height[left-1] > leftMax {
				leftMax = height[left-1]
			}
			// 当前柱子比左边最高的柱子还低才能接水
			diff := leftMax - height[left]
			if diff > 0 {
				water += diff
			}
			// 接完后往中间靠
			left++
		}
	}

	return water
}

// 解法 1，通过俩数组记录每一根柱子左右两边的情况
func trap1(height []int) int {
	if height == nil || len(height) < 3 {
		return 0
	}

	l := len(height)

	// leftDp(i) 代表第 i 根柱子左边最高的柱子
	leftDp := make([]int, l)
	leftDp[0] = height[0]
	for i := 1; i < l-1; i++ {
		leftDp[i] = leftDp[i-1]
		if leftDp[i] < height[i] {
			leftDp[i] = height[i]
		}
	}

	// rightDp(i) 代表第 i 根柱子右边最高的柱子
	rightDp := make([]int, l)
	rightDp[l-1] = height[l-1]
	for i := l - 2; i >= 1; i-- {
		rightDp[i] = rightDp[i+1]
		if rightDp[i] < height[i] {
			rightDp[i] = height[i]
		}
	}

	// 雨水
	water := 0

	// 从第 2 根柱子开始到倒数第 2 根柱子结束，能接多少雨水
	for i := 1; i < l-1; i++ {
		// 先找出左右俩边最低的柱子
		lower := leftDp[i]
		if lower > rightDp[i] {
			lower = rightDp[i]
		}

		// 如果当前柱子比俩边最低都还要低，那肯定接不了雨水了
		if height[i] >= lower {
			continue
		}

		// 来到这里，肯定可以接雨水，注意需要减去当前柱子的高度
		water += lower - height[i]
	}

	return water
}
