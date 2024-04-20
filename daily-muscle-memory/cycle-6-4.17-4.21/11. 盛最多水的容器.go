// @Author: Ciusyan 4/19/24

package cycle_6_4_17_4_21

// https://leetcode.cn/problems/container-with-most-water/

/*
*
思路重复：
对于柱子间能装多少水，其实就是求出最大面积
那么准备双指针，分别指向左右。
然后每次都取最低的那个作为高。
底就是 右-左
最终面积就是：底 * 高
那么有没有比原先的面积大呢？比一下就好，大就更新结果值。
然后往中间靠，但是在往中间靠的时候，可以比较一下，是不是 <= lower，如果是，就跳过不用算了
因为就算计算，高和底都比上一次的小，结果也会比上一次的小，就没必要算了。
*/
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
