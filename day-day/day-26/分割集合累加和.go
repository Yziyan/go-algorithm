// @Author: Ciusyan 12/8/23

package day_26

/**
给定一个正数数组 arr，
请把 arr 中所有的数分成两个集合，尽量让两个集合的累加和接近
返回最接近的情况下，较小集合的累加和
*/

func splitSumClosed(arr []int) int {
	if arr == nil || len(arr) < 2 {
		return 0
	}

	// 先来憋一个暴力递归，含义是：
	// 当前处于 arr[cur] 元素，使用 [cur ...] 的元素，
	// 返回最接近 remain 的最小累加和
	var process func(arr []int, cur int, remain int) int
	process = func(arr []int, cur int, remain int) int {
		if cur == len(arr) {
			// 说明没有元素可以累加出 remain 了，那么最接近 remain 的也只能是 0 了
			return 0
		}

		// 来到这里，列举所有可能：
		// 1.不选择当前的元素
		p1 := process(arr, cur+1, remain)
		// 2.选当前元素，但是要保证选择后不能超过 remain
		p2 := 0
		if arr[cur] <= remain {
			// 说明可选，选择了当前的元素，
			// 那么剩下 [cur+1 ...] 个元素，就只需要累加出接近 remain-arr[cur] 的和
			p2 = arr[cur] + process(arr, cur+1, remain-arr[cur])
		}

		// 那么结论就是：两种累加和的最大值因为他们越大越接近 remain
		return max(p1, p2)
	}

	// 先求出所有元素的累加和
	sum := 0
	for _, v := range arr {
		sum += v
	}

	// 那么主函数怎么调用呢？
	// 既然要拆分，而且让两个集合的累加和尽量接近，所以最平均的就是接近数组累加和的一半
	return process(arr, 0, sum>>1)
}
