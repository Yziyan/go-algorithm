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

// 思路重复 + 代码重复
/*
1.核心就是：一根柱子一根柱子的计算，即竖着看能接的雨水量
2.能接雨水的柱子，它的左右两边一定有比它高的柱子。
3.所以准备两个指针，left 和 right，每次都找出他们最低的一边 lower。
4.那么就只用考虑 lower 所在边即可，另一遍一定比它高
5.所以又准备了一个不断更新的 lowerMax，这个只可能被每一次 lower 所在边更新。
	5.1 要么原本的值本来就在 lower 所在边，因为每次只有较低的柱子才会往中心挪动
	5.2 要么此次会使用 lower 更新 lowerMax，那么也会变成 lower 所在边
6. 所以此次能接住的雨水，就是 lowerMax-lower

// 伪代码
1.判空
2.准备 water 和 lowerMax，还有双指针 left 和 right
3.进行遍历，直至两个指针相遇
	3.1 找出 lower，并且低的一边，要往中心靠
	3.2 尝试更新 lowerMax
	3.3 添加上此次接的雨水量
4.返回water
*/
func trap2(height []int) int {
	if height == nil || len(height) == 0 {
		return 0
	}

	var (
		water    = 0
		lowerMax = 0

		left  = 0
		right = len(height) - 1
	)

	for left < right {
		lower := height[left]
		left++
		if height[right] < lower {
			lower = height[right]
			right--
			left--
		}

		if lower > lowerMax {
			lowerMax = lower
		}

		water += lowerMax - lower
	}

	return water
}
