// @Author: Ciusyan 4/14/24

package cycle_5_4_12_4_16

// https://leetcode.cn/problems/spiral-matrix/

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	var (
		rowL = len(matrix)
		colL = len(matrix[0])
		res  = make([]int, 0, rowL*colL)

		// 准备四个指针，分别代表上下左右
		top    = 0
		bottom = rowL - 1
		left   = 0
		right  = colL - 1
	)

	// 左右都不能越界
	for left <= right && top <= bottom {
		// 左上 -> 右上
		for cur := left; cur <= right; cur++ {
			res = append(res, matrix[top][cur])
		}
		top++

		if top > bottom {
			// 说明收集完毕了
			break
		}

		// 右上 -> 右下
		for cur := top; cur <= bottom; cur++ {
			res = append(res, matrix[cur][right])
		}
		right--

		if left > right {
			// 说明收集完毕了
			break
		}

		// 右下 -> 左下
		for cur := right; cur >= left; cur-- {
			res = append(res, matrix[bottom][cur])
		}
		bottom--

		// 左下 -> 左上
		for cur := bottom; cur >= top; cur-- {
			res = append(res, matrix[cur][left])
		}
		left++
	}

	return res
}
