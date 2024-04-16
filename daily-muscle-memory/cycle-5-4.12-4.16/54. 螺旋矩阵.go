// @Author: Ciusyan 4/14/24

package cycle_5_4_12_4_16

// https://leetcode.cn/problems/spiral-matrix/

/**
思路重复：
需要螺旋收集矩阵的值，那么我们准备（上、下、左、右）四个指针，按想要的规则模拟即可：
只要（上下）（左右）都不越界。就说明还没有收集完毕。那么按照顺序，依次收集：
左上 --> 右上 -> 右下 -> 左下 -> 左上
直至所有元素都收集完毕，但是其中要有一个注意点，就是：
每收集完一行 or 一列后，会将相应行或者相应列变更，变更后，要保证不会越界。否则说明全部收集完毕了。
*/

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}

	var (
		rowL = len(matrix)
		colL = len(matrix[0])
		res  = make([]int, 0, rowL*colL)

		top, bottom = 0, rowL - 1
		left, right = 0, colL - 1
	)

	for top <= bottom && left <= right {
		// 左上 -> 右上
		for cur := left; cur <= right; cur++ {
			res = append(res, matrix[top][cur])
		}
		top++
		if top > bottom {
			break
		}
		// 右上 -> 右下
		for cur := top; cur <= bottom; cur++ {
			res = append(res, matrix[cur][right])
		}
		right--
		if left > right {
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

func spiralOrder1(matrix [][]int) []int {
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
