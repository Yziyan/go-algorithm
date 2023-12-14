// @Author: Ciusyan 12/14/23

package day_27

/**
给定一个整型数组 arr，和一个整数 num
某个 arr 中的子数组 sub，如果想达标，必须满足：sub 中最大值 – sub 中最小值 <= num，
返回 arr 中达标子数组的数量
*/

func allLessNumSubArray(arr []int, sum int) int {
	if arr == nil || len(arr) == 0 || sum < 0 {
		return 0
	}

	n := len(arr)

	// 准备两个双端队列，分别用于更新窗口内最大值和最小值，队列里放索引
	queueMax := NewDoubleQueue()
	queueMin := NewDoubleQueue()

	// 待求结果
	count := 0
	// 因为左右边界都不回退
	r := 0
	for l := 0; l < n; l++ {
		// 依次求解：[L ... R] [L+1 ... R] 满足条件的解
		// 但再怎么往后跑，r 也不能越界，
		for r < n {
			// 先更新窗口内的最大值和最小值
			for queueMax.Size() != 0 && arr[r] >= arr[queueMax.GetRight()] {
				// 说明添加 arr[r] 有障碍，先清除掉，要保证 queueMax 从大到小
				queueMax.PollRight()
			}
			// 可以放心添加就进入 queueMax 了
			queueMax.OfferRight(r)
			for queueMin.Size() != 0 && arr[r] <= arr[queueMin.GetRight()] {
				// 说明添加 arr[r] 有障碍，先清除掉，要保证 queueMin 从小到大
				queueMin.PollRight()
			}
			// 可以放心添加就进入 queueMin 了
			queueMin.OfferRight(r)

			// 还得看看满不满足题意，即 max - min  <= sum
			if arr[queueMax.GetLeft()]-arr[queueMin.GetLeft()] > sum {
				// 说明不满足条件了，去收集 [l ... r] 中，以 l 打头的子数组
				break
			}

			// 说明还可能往后扩张
			r++
		}

		// 来到这里，说明找到了第一个 [L ... R] 上，第一个不满足条件的边界
		count += r - l
		// 退出前，窗口左边界要前进了，需要检查窗口最大值 or 最小值有没有失效。
		if l == queueMax.GetLeft() {
			// 说明下一次更新，就失效了
			queueMax.PollLeft()
		}
		if l == queueMin.GetLeft() {
			// 说明下一次更新，就失效了
			queueMin.PollLeft()
		}
	}

	return count
}
