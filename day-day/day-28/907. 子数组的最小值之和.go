// @Author: Ciusyan 12/22/23

package day_28

// https://leetcode.cn/problems/sum-of-subarray-minimums/

func sumSubarrayMins(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return 0
	}

	res := 0
	// 我们转换一下标准，即：若以 arr[i] 位置作为最小值，有多少个子数组，那么这些累加和就是 arr[i] * 数量
	// 那么以所有位置作为最小值，求出来的累加和相加，就是最终的答案

	// 先求出左右两边能扩张的范围
	extant := getIdxExtant(arr)
	// 挨个作为最小值，求解一个结果
	for i := range arr {
		// 有多少个以 arr[i] 作为最小值的子数组呢？有 左边 * 右边
		left := i - extant[i][0]
		right := extant[i][1] - i

		// 那么当前的一个累加和就是
		res += left * right * arr[i]
	}

	// 答案要求模上 10^9 + 7
	return res % 1000000007
}

// 返回值：
// res[i][0] 左边第一个小于 arr[i] 的位置，
// res[i][1] 右边第一个小于等于 arr[i] 的位置
func getIdxExtant(arr []int) [][2]int {
	l := len(arr)
	res := make([][2]int, l)
	// 准备一个单调栈，使用数组模拟
	stack := make([]int, l)
	si := -1 // 代表栈顶索引，-1 代表栈为空

	// 挨个值加入栈里面，挨个求解
	for i := range arr {
		// 在加入之前，需要查看加入后是否会违反单调性
		for si != -1 && arr[i] <= arr[stack[si]] {
			// 说明当前元素比栈顶还小，该弹出求解结果了
			popIdx := stack[si]
			si--
			leftLessIdx := -1
			if si != -1 {
				// 说明弹出栈顶后还有值
				leftLessIdx = stack[si]
			}

			// 设置答案
			res[popIdx] = [2]int{leftLessIdx, i}
		}
		// 到这里，说明可以安全加入栈了，只加索引
		si++
		stack[si] = i
	}

	// 说明元素加完了，但是还需要结算栈里的元素
	for si != -1 {
		// 弹出栈顶结算，自然弹出
		popIdx := stack[si]
		si--

		leftLessIdx := -1
		if si != -1 {
			// 说明弹出后还有元素，
			leftLessIdx = stack[si]
		}

		// 设置答案，右边是自然弹出的，所以能到最右边
		res[popIdx] = [2]int{leftLessIdx, l}
	}

	return res
}
