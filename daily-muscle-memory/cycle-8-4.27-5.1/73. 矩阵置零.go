// @Author: Ciusyan 4/29/24

package cycle_8_4_27_5_1

// https://leetcode.cn/problems/set-matrix-zeroes/

func setZeroes(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}

	rowL := len(matrix)
	colL := len(matrix[0])

	firstRow0 := false
	firstCol0 := false

	// 找到第一行是否为 0
	for col := 0; col < colL; col++ {
		if matrix[0][col] == 0 {
			firstRow0 = true
		}
	}
	// 找到第一列是否为 0
	for row := 0; row < rowL; row++ {
		if matrix[row][0] == 0 {
			firstCol0 = true
		}
	}

	for row := 1; row < rowL; row++ {
		for col := 1; col < colL; col++ {
			if matrix[row][col] == 0 {
				matrix[row][0] = 0
				matrix[0][col] = 0
			}
		}
	}

	// 挨行检查，需要置零的行
	for row := 1; row < rowL; row++ {
		if matrix[row][0] == 0 {
			for col := 1; col < colL; col++ {
				matrix[row][col] = 0
			}
		}
	}

	// 挨列检查，需要置零的列
	for col := 1; col < colL; col++ {
		if matrix[0][col] == 0 {
			for row := 1; row < rowL; row++ {
				matrix[row][col] = 0
			}
		}
	}

	// 第一行是否要置零
	if firstRow0 {
		for col := 0; col < colL; col++ {
			matrix[0][col] = 0
		}
	}
	// 第一列是否要置零
	if firstCol0 {
		for row := 0; row < rowL; row++ {
			matrix[row][0] = 0
		}
	}
}
