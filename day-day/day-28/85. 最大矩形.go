// @Author: Ciusyan 12/20/23

package day_28

// https://leetcode.cn/problems/maximal-rectangle/description/

func maximalRectangle(matrix [][]byte) int {
	if matrix == nil || len(matrix) == 0 {
		return 0
	}

	// 我们如果挨行求解下来，以每一行作为最后一行，去求解一个最大的结果。
	// 到下一行的时候，将上一行的值累加下来，但是要保证合理，
	// 那么就能转换成求解直方图最大矩形的面积问题了
	colLen := len(matrix[0])

	res := 0
	rowHeights := make([]int, colLen)
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < colLen; col++ {
			if matrix[row][col] == '0' {
				// 说明当前行 col 不联通了
				rowHeights[col] = 0
			} else {
				// 说明还联通，以前的值加上当前行的值
				rowHeights[col] += 1
			}
		}
		// 挨行为底，去求解直方图的面积
		res = max(res, perRowMaxArea(rowHeights))
	}

	return res
}

// 求解以每一行为底构成的直方图，最大矩形的面积
func perRowMaxArea(heights []int) int {
	res := 0
	n := len(heights)
	// 准备一个单调栈，求解每一个位置固定为高时，能求解出来的最大面积
	stack := NewStack()
	// 挨个高先加入单调栈里
	for i := range heights {
		// 但是加入当前高，不能违反单调性
		for stack.Size() != 0 && heights[i] <= heights[stack.Peek()] {
			// 说明当前元素比栈顶还小，那么栈顶找到了右边第一个比自己还矮的高，
			// 可以弹出它来求解结果了
			popIdx := stack.Pop()
			leftLessIdx := -1
			if stack.Size() != 0 {
				// 说明弹出栈顶后还有元素
				leftLessIdx = stack.Peek()
			}
			// 面积就是：底 * 高
			area := (i - leftLessIdx - 1) * heights[popIdx]
			res = max(res, area)
		}
		// 到达这里，说明可以顺利加入栈了，记得只加入索引即可
		stack.Push(i)
	}

	// 来到这里说明所有元素都加完了，还需要将栈清空，
	for stack.Size() != 0 {
		// 弹出栈顶计算结果
		popIdx := stack.Pop()
		leftLessIdx := -1
		if stack.Size() != 0 {
			// 说明弹出栈顶后还不为空
			leftLessIdx = stack.Peek()
		}
		// 当前 popIdx 作为高时，能求解出来的面积
		area := (n - leftLessIdx - 1) * heights[popIdx]
		res = max(res, area)
	}

	return res
}
