// @Author: Ciusyan 12/19/23

package day_28

// https://leetcode.cn/problems/largest-rectangle-in-histogram/description/

func largestRectangleArea(heights []int) int {
	if heights == nil || len(heights) == 0 {
		return 0
	}

	n := len(heights)
	res := 0
	// 准备一个单调栈，存放索引
	stack := NewStack()

	// 挨个以 heights[i] 固定为高，看看左右两边能扩张多远。
	// 当 i 号柱子作为高时，它的左右扩张到最远必须不能比自己矮，否则不能当做矩形
	// 所以相当于寻找，i 号柱子左右两边第一根比自己还矮的柱子
	for i := range heights {
		// 尝试将 i 号柱子放入单调栈中
		for stack.Size() != 0 && heights[i] <= heights[stack.Peek()] {
			// 说明当前栈顶的柱子比较矮，需要先清空现场，记录结果
			popIdx := stack.Pop()
			leftLessIdx := -1
			if stack.Size() != 0 {
				// 说明栈顶弹出后还有元素
				leftLessIdx = stack.Peek()
			}
			// 当前面积，固定 popIdx 为高，能扩张到最远的长度为底
			cur := heights[popIdx] * (i - leftLessIdx - 1)

			// 求解结果
			res = max(res, cur)
		}
		// 将当前柱子入栈，稍后需要弹出来求解这个柱子为高时的最大面积
		stack.Push(i)
	}

	// 当所有柱子都遍历完后，需要清空栈，
	for stack.Size() != 0 {
		// 弹出柱子作为高，求解
		popIdx := stack.Pop()
		leftLessIdx := -1
		if stack.Size() != 0 {
			// 栈顶弹出后还有元素
			leftLessIdx = stack.Peek()
		}

		// 求出当前面积：高 * 底
		cur := heights[popIdx] * (n - leftLessIdx - 1)
		// 记录结果
		res = max(res, cur)
	}

	return res
}
