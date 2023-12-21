// @Author: Ciusyan 12/21/23

package day_28

// https://leetcode.cn/problems/count-submatrices-with-all-ones/

func numSubmat(mat [][]int) int {
	if mat == nil || len(mat) == 0 || len(mat[0]) == 0 {
		return 0
	}

	colLen := len(mat[0])
	res := 0
	// 以每一行作为底，去求解能够有多少矩形，所有行求解的结果相加，就是答案了
	// 代表以挨行作为底的直方图
	rowHeights := make([]int, colLen)
	for row := 0; row < len(mat); row++ {
		for col := 0; col < colLen; col++ {
			if mat[row][col] == 1 {
				// 说明当前列还联通，能够贡献 1 的高度
				rowHeights[col]++
			} else {
				// 当 (row, col) 不联通了，前面再高也得 ban 掉
				rowHeights[col] = 0
			}
		}

		// 挨行去求解结果。
		res += getRowHeightsNum(rowHeights)
	}

	return res
}

// 求解直方图 heights，能够围出多少矩形
func getRowHeightsNum(heights []int) int {

	// 那么现在得挨列求解了，每一列计算能够贡献多少矩形，那么所有列相加就是答案
	res := 0
	l := len(heights)
	// 准备一个单调栈，只存放索引
	stack := NewStack()
	// 现将元素挨个加入栈里
	for i := range heights {
		// 但是加入不能违反单调性
		for stack.Size() != 0 && heights[i] <= heights[stack.Peek()] {
			// 说明当前元素更矮，加入后会违反栈的单调性，
			// 弹出栈顶计算结果
			popIdx := stack.Pop()
			if heights[i] == heights[popIdx] {
				// 说明栈顶和当前元素相等，那么等到后面相等的那个再算
				continue
			}
			// 来到这里，可以计算弹出柱子的结果了
			leftLessIdx := -1
			if stack.Size() != 0 {
				// 说明弹出栈顶后还有元素
				leftLessIdx = stack.Peek()
			}
			// 要在此柱子，计算高度位于 popIdx 和 max(leftIdx, rightIdx) 高度之间的矩形，
			// 每一个高都要算，而且他们公司都一样，所以就是看需要计算多少次
			count := heights[popIdx] - heights[i] // 默认左边没有柱子了，以右边为准
			if leftLessIdx != -1 {
				// 说明也要参考左边，取较高的为准，
				count = heights[popIdx] - max(heights[leftLessIdx], heights[i])
			}

			// 那么以 [popIdx, popIdx-max(leftIdx, rightIdx)] 为高的，都有多少个矩形呢？
			// 成一个等差数列分布，所以是等差数列求和的公式：num = n*(n+1)/2
			n := i - leftLessIdx - 1
			res += count * (n * (n + 1) >> 1)
		}
		// 来到这里，可以放心加入栈里等待计算结果了，但是放的是索引
		stack.Push(i)
	}

	// 到这里说明所有元素都被加入栈了，还需要清算结果
	for stack.Size() != 0 {
		// 也是弹出栈顶计算结论
		popIdx := stack.Pop()
		leftLessIdx := -1
		if stack.Size() != 0 {
			// 说明栈顶弹出了还有元素
			leftLessIdx = stack.Peek()
		}

		// 还是需要计算要求解多少次
		count := heights[popIdx] // 因为是自然弹出的，右边不需要参考了，默认 popIdx-0
		if leftLessIdx != -1 {
			// 说明左边还需要参考
			count = heights[popIdx] - heights[leftLessIdx]
		}
		n := l - leftLessIdx - 1
		res += count * (n * (n + 1) >> 1)
	}

	return res
}
