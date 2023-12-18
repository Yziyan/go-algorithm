// @Author: Ciusyan 12/18/23

package day_28

import "math"

/**
给定一个只包含正数的数组 arr，arr 中任何一个子数组 sub，
一定都可以算出 (sub累加和 )* (sub中的最小值) 是什么，
那么所有子数组中，这个值最大是多少？
*/

func allTimesMinToMax(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return 0
	}

	n := len(arr)
	// 方便之后快速求解子数组的累加和，现将其全部累加起来
	sum := make([]int, n)
	sum[0] = arr[0]
	for i := 1; i < n; i++ {
		sum[i] = sum[i-1] + arr[i]
	}
	// 累加后，要求 i...j 单独的累加和，那么就是 sum[j] - sum[i-1]

	res := math.MinInt
	// 准备一个单调栈，里面存放索引
	stack := NewStack()
	// 遍历所有元素
	for i := range arr {
		// 想要添加到栈里面，但是不能违背单调性
		for stack.Size() != 0 && arr[i] <= arr[stack.Peek()] {
			// 说明当前元素小于等于栈顶，需要弹出清算答案。
			// 为什么等于的时候需要弹出呢？虽然等于的时候弹出来算答案了，答案是错的，
			// 但是我们最终会将等于的这一个值，也加入栈里面，后加入的相等的值，会纠正前一个弹出的值的答案。
			popIdx := stack.Pop()
			// 弹出索引范围的累加和，默认是从最左到最右，那么就是累加到最后一个位置的和
			left2rightSum := sum[i-1]
			if stack.Size() != 0 {
				// 说明弹出后栈还有元素，压着的肯定比当前 popIdx 要小，所以要截止，减去之前的累加和
				left2rightSum -= sum[stack.Peek()]
			}

			// 计算结果
			res = max(res, left2rightSum*arr[popIdx])
		}
		// 到达这里，说明可以正常加入了
		stack.Push(i)
	}

	// 元素遍历完后，还得清空栈的元素，设置结果
	for stack.Size() != 0 {
		// 也是弹出栈顶，求一个结果
		popIdx := stack.Pop()
		// 弹出索引范围的累加和，默认是从最左到最右，那么就是累加到最后一个位置的和
		left2rightSum := sum[n-1]
		if stack.Size() != 0 {
			// 说明弹出后栈还有元素，压着的肯定比当前 popIdx 要小，所以要截止
			left2rightSum -= sum[stack.Peek()]
		}

		// 求解结果
		res = max(res, left2rightSum*arr[popIdx])
	}

	return res
}
