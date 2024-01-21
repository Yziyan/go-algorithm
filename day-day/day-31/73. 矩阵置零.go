// @Author: Ciusyan 1/21/24

package day_31

// https://leetcode.cn/problems/set-matrix-zeroes/

func setZeroes(matrix [][]int) {
	if matrix == nil || len(matrix) == 0 {
		return
	}

	var (
		rowL = len(matrix)
		colL = len(matrix[0])

		// 用两个变量标记 0行和0列 是否需要置为 0
		row0 = false
		col0 = false
	)

	for row := 0; row < rowL; row++ {
		if matrix[row][0] == 0 {
			// 代表第 0列 有 0
			col0 = true
		}
	}

	for col := 0; col < colL; col++ {
		if matrix[0][col] == 0 {
			// 代表第 0行 有 0
			row0 = true
		}
	}

	// 遍历剩下的数字，用原始数组第一行和第一列分别标识对应行和对应列是否需要置为 0，
	for row := 1; row < rowL; row++ {
		for col := 1; col < colL; col++ {
			if matrix[row][col] == 0 {
				// 说明对应行和对应列最后都需要置为零，先用 0 打个标记
				matrix[row][0] = 0
				matrix[0][col] = 0
			}
		}
	}

	// 看看哪些行需要置为零
	for row := 1; row < rowL; row++ {
		if matrix[row][0] != 0 {
			// 说明这行不需要置零
			continue
		}
		// 需要将这一行置零
		for col := 0; col < colL; col++ {
			matrix[row][col] = 0
		}
	}

	// 看看哪些列需要置为零
	for col := 1; col < colL; col++ {
		if matrix[0][col] != 0 {
			// 说明这一列不需要置零
			continue
		}
		// 需要将这一列置零
		for row := 0; row < rowL; row++ {
			matrix[row][col] = 0
		}
	}

	// 最后再看看第 0行和0列 需要置零吗
	if row0 {
		// 说明需要将第 0 行置零
		for col := 0; col < colL; col++ {
			matrix[0][col] = 0
		}
	}

	if col0 {
		// 说明需要将地 0 列置零
		for row := 0; row < rowL; row++ {
			matrix[row][0] = 0
		}
	}
}
