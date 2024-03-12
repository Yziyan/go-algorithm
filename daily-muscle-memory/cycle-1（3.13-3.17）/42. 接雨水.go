// @Author: Ciusyan 3/13/24

package cycle_1_3_13_3_17_

// https://leetcode.cn/problems/trapping-rain-water/

func trap(height []int) int {
	if height == nil || len(height) < 3 {
		return 0
	}

	var (
		water    = 0 // 最终能接的雨水
		lowerMax = 0 // 首先是每次都找最低的一边，然后和之前较低的比较，较大的那个

		left  = 0
		right = len(height) - 1
	)

	// 当两个指针相遇的时候，就接完雨水了
	for left < right {
		// 找出两根柱子较低的一根，并且谁低谁往后移动
		lower := height[left] // 先假设 left 要低
		left++
		if height[right] < lower {
			// 说明 right 更低
			lower = height[right]
			right--
			left-- // 还原上面默认左边最低移动的指针
		}

		// 上面找到了这一次的 lower，看看能不能比 lowerMax 还大
		if lower > lowerMax {
			// 说明比上一次最低的柱子还高，更新 lowerMax
			lowerMax = lower
		}

		// 最后更新雨水量，
		// 因为上面找到的 lower 所在的边，就确定了这边一定比另一边低， 可以放心的看自己这一边，
		// 又由于低的一边才会往中间挪动，所以 lowerMax 一定是利用自己所在边更新的，
		// 那么就可以直接看自己所在边和最低柱子的高度差
		water += lowerMax - lower
	}

	return water
}
